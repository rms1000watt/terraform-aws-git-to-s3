module "serverless" {
  source  = "rms1000watt/serverless-tf/aws"
  version = "0.2.10"

  functions = [
    {
      file       = "main.go"
      vendor_cmd = "govendor sync"
      role_arn   = "arn:aws:iam::${local.account_id}:role/${local.role_name}"
      http_path  = "update"

      env_keys = "GIT_USER GIT_PASS GIT_URL"
      env_vals = "${var.git_user} ${var.git_pass} ${var.git_url}"
    },
  ]
}
