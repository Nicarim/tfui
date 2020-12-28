resource "random_string" "random" {
  length = 16
  special = true
  override_special = "/@£$"
}
resource "random_string" "_one" {
  length = 16
  special = true
  override_special = "/@£$"
}
resource "random_string" "_two" {
  length = 16
  special = true
  override_special = "/@£$"
}
resource "random_string" "_three" {
  length = 16
  special = true
  override_special = "/@£$"
}
