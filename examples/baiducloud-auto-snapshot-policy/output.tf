output "autosnapshotpolicy" {
  value = "${baiducloud_auto_snapshot_policy.my-asp}"
}

output "auto_snapshot_policies" {
  value = "${data.baiducloud_auto_snapshot_policies.default.auto_snapshot_policies}"
}