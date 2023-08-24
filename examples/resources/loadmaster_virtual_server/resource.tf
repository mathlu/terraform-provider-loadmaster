resource "loadmaster_virtual_service" "foo" {
  address  = "192.168.1.10"
  protocol = "tcp"
  port     = "8080"
  nickname = "bar"
  type     = "gen"
}
