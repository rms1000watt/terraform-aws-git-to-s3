resource "aws_s3_bucket" "0" {
  bucket = "git-to-s3-${local.project_name}-${local.account_id}"
  acl    = "private"
}
