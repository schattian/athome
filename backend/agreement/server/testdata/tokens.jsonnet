local jwt = import './jwt.libsonnet';
{
  local uid = 4,
  valid: {
    user_id: uid,
    sign_token: jwt.sign.valid,
    access_token: jwt.auth.valid,
    refresh_token: jwt.auth.valid,
  },
  expired_access: $.valid {
    access_token: jwt.auth.expired,
  },
  expired_refresh: $.valid {
    refresh_token: jwt.auth.expired,
  },
  expired_sign: $.valid {
    sign_token: jwt.sign.expired,
  },
  expired: {
    user_id: uid,
    sign_token: jwt.sign.expired,
    access_token: jwt.auth.expired,
    refresh_token: jwt.auth.expired,
  },
}
