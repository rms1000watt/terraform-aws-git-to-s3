locals {
  path_root_list = "${split("/", path.root)}"
  path_root_len  = "${length(local.path_root_list)}"
  parent_dir     = "${element(local.path_root_list, local.path_root_len - 1)}"

  region     = "${data.aws_region.0.name}"
  account_id = "${data.aws_caller_identity.0.account_id}"
}

locals {
  project_name = "${var.project_name != "" ? var.project_name : local.parent_dir}"
}
