resource "aws_route53_record" "example_com" {
  zone_id = "ZONEID"
  name    = "example.com."
  type    = "MX"
  ttl     = "300"
  records = [
    "10 mail.example.com"
  ]
}

resource "aws_route53_record" "www_example_com" {
  zone_id = "ZONEID"
  name    = "www.example.com."
  type    = "CNAME"
  ttl     = "300"
  records = [
    "www.ap-southeast-1.elb.amazonaws.com"
  ]
}

resource "aws_route53_record" "cloudflare-verify_example_com" {
  zone_id = "ZONEID"
  name    = "cloudflare-verify.example.com."
  type    = "TXT"
  ttl     = "300"
  records = [
    "qwertyuiop-asdfghjkl"
  ]
}

resource "aws_route53_record" "web_example_com" {
  zone_id = "ZONEID"
  name    = "web.example.com."
  type    = "A"
  ttl     = "300"
  records = [
    "111.111.111.111"
  ]
}

resource "aws_route53_record" "app_example_com" {
  zone_id = "ZONEID"
  name    = "app.example.com."
  type    = "A"

  alias {
    name                   = "app.prod-utility.ap-southeast-1.elb.amazonaws.com."
    zone_id                = "QWERTYUIOP"
    evaluate_target_health = "false"
  }
}

