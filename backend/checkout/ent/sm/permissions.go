package sm

func (s *State) userPermissions(entity Stateful, uid uint64) (p permissions) {
	switch uid {
	case entity.GetConsumerId():
		p = s.consumer
	case entity.GetMerchantId():
		p = s.merchant
	case entity.GetServiceProviderId():
		p = s.serviceProvider
	}

	return
}

var (
	all = permissions{cancellable: true, prevable: true, nextable: true}

	onlyPrev   = permissions{prevable: true}
	onlyNext   = permissions{nextable: true}
	onlyCancel = permissions{cancellable: true}
)
