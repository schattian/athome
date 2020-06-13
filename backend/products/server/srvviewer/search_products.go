package srvviewer

import (
	"context"
	"encoding/base64"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/athomecomar/athome/backend/products/ent"
	"github.com/athomecomar/athome/backend/products/server"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"google.golang.org/grpc/status"
)

func (s *Server) SearchProducts(ctx context.Context, in *pbproducts.SearchProductsRequest) (*pbproducts.SearchProductsResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sem, semCloser, err := pbconf.ConnSemantic(ctx)
	if err != nil {
		return nil, err
	}
	defer semCloser()

	img, imgCloser, err := pbconf.ConnImages(ctx)
	if err != nil {
		return nil, err
	}
	defer imgCloser()

	users, usersCloser, err := pbconf.ConnUsers(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	return s.searchProducts(ctx, db, sem, users, img, in)
}

func (s *Server) searchProducts(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, users pbusers.ViewerClient, img pbimages.ImagesClient, in *pbproducts.SearchProductsRequest) (*pbproducts.SearchProductsResponse, error) {
	q, err := toNormal(in.GetQuery())
	if err != nil {
		return nil, err
	}
	page := in.GetPage()
	cursorId, err := b64DecodeId(page.GetCursor())
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryxContext(ctx, `SELECT id, title, price FROM products 
    WHERE id < $1
    AND lower(unaccent(title)) ILIKE ESCAPE $2
    ORDER BY id LIMIT $3`, cursorId, q, page.GetSize())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	var prods []*ent.Product
	for rows.Next() {
		pr := &ent.Product{}
		err = rows.StructScan(pr)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}
		prods = append(prods, pr)
	}

	resp := &pbproducts.SearchProductsResponse{Page: &pbproducts.PageResponse{}}
	count, err := db.QueryxContext(ctx, `SELECT COUNT(*) FROM products WHERE lower(unaccent(title)) ILIKE ESCAPE $1`, q)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	for count.Next() {
		err := count.Scan(&resp.Page.TotalSize)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "Scan: %v", err)
		}
	}

	if len(prods) > 0 {
		resp.Page.NextCursor = b64EncodeId(prods[len(prods)].Id)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})

	var lock sync.RWMutex
	for _, pr := range prods {
		wg.Add(1)
		pr := pr
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := pr.ToPbSearchResult(ctx, users, img)
			if err != nil {
				errCh <- status.Errorf(xerrors.Internal, "ToPbSearchResult: %v", err)
			}
			lock.Lock()
			defer lock.Unlock()
			resp.Products[pr.Id] = res
		}(&wg)
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case err := <-errCh:
			return nil, err
		case <-done:
			return resp, nil
		}
	}
}

func toNormal(s string) (string, error) {
	s = strings.ToLower(s)
	s, err := stripAccents(s)
	if err != nil {
		return "", err
	}
	s = removeNonWord(s)
	s = strings.TrimSpace(s)
	return s, nil
}

func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func stripAccents(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, s)
	if err != nil {
		return "", status.Errorf(xerrors.Internal, "transform.String: %v", err)
	}
	return result, nil
}

func removeNonWord(s string) string {
	return clean([]byte(s))
}

func b64DecodeId(str string) (uint64, error) {
	bytes, err := base64.RawStdEncoding.DecodeString(str)
	if err != nil {
		return 0, status.Errorf(xerrors.InvalidArgument, "base64.DecodeString: %v", err)
	}
	id, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0, status.Errorf(xerrors.InvalidArgument, "strconv.Atoi: %v", err)
	}
	return uint64(id), nil
}

func b64EncodeId(uid uint64) string {
	return base64.RawStdEncoding.EncodeToString([]byte(strconv.Itoa(int(uid))))
}
