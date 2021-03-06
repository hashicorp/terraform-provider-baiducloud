package baiducloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const (
	testAccCertsDataSourceName          = "data.baiducloud_certs.default"
	testAccCertsDataSourceAttrKeyPrefix = "certs.0."
)

//lintignore:AT003
func TestAccBaiduCloudCertDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,

		Steps: []resource.TestStep{
			{
				Config: testAccCertDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaiduCloudDataSourceId(testAccCertsDataSourceName),
					resource.TestCheckResourceAttr(testAccCertsDataSourceName, "certs.#", "1"),
					resource.TestCheckResourceAttr(testAccCertsDataSourceName, testAccCertsDataSourceAttrKeyPrefix+"cert_name", BaiduCloudTestResourceAttrNamePrefix+"Cert"),
					resource.TestCheckResourceAttr(testAccCertsDataSourceName, testAccCertsDataSourceAttrKeyPrefix+"cert_type", "1"),
					resource.TestCheckResourceAttrSet(testAccCertsDataSourceName, testAccCertsDataSourceAttrKeyPrefix+"cert_start_time"),
					resource.TestCheckResourceAttrSet(testAccCertsDataSourceName, testAccCertsDataSourceAttrKeyPrefix+"cert_stop_time"),
					resource.TestCheckResourceAttrSet(testAccCertsDataSourceName, testAccCertsDataSourceAttrKeyPrefix+"cert_create_time"),
					resource.TestCheckResourceAttrSet(testAccCertsDataSourceName, testAccCertsDataSourceAttrKeyPrefix+"cert_update_time"),
				),
			},
		},
	})
}

func testAccCertDataSourceConfig() string {
	return fmt.Sprintf(`
resource "baiducloud_cert" "default2" {
  cert_name         = "%s"
  cert_server_data  = "-----BEGIN CERTIFICATE-----\nMIIEGzCCA8CgAwIBAgIQBHVIJNCDJKsC1maaUVgqdjAKBggqhkjOPQQDAjByMQswCQYDVQQGEwJDTjElMCMGA1UEChMcVHJ1c3RBc2lhIFRlY2hub2xvZ2llcywgSW5jLjEdMBsGA1UECxMURG9tYWluIFZhbGlkYXRlZCBTU0wxHTAbBgNVBAMTFFRydXN0QXNpYSBUTFMgRUNDIENBMB4XDTE5MDkwNjAwMDAwMFoXDTIwMDkwNTEyMDAwMFowHzEdMBsGA1UEAxMUdGVzdC55aW5jaGVuZ2ZlbmcuY24wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR+aGvOdizh+oAWwT6829WdcZw7oBJVU1UvKQdm7dW/7SIdrMEWq6NIWaERMKkLD6gQ6Y5KFV9oDQdSocGBtBvLo4ICiTCCAoUwHwYDVR0jBBgwFoAUEoZEZiYIVCaPZTeyKU4mIeCTvtswHQYDVR0OBBYEFAichc0eFh+KdwMYjD7Pbvc8Q80IMB8GA1UdEQQYMBaCFHRlc3QueWluY2hlbmdmZW5nLmNuMA4GA1UdDwEB/wQEAwIHgDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwTAYDVR0gBEUwQzA3BglghkgBhv1sAQIwKjAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGlnaWNlcnQuY29tL0NQUzAIBgZngQwBAgEwgZIGCCsGAQUFBwEBBIGFMIGCMDQGCCsGAQUFBzABhihodHRwOi8vc3RhdHVzZi5kaWdpdGFsY2VydHZhbGlkYXRpb24uY29tMEoGCCsGAQUFBzAChj5odHRwOi8vY2FjZXJ0cy5kaWdpdGFsY2VydHZhbGlkYXRpb24uY29tL1RydXN0QXNpYVRMU0VDQ0NBLmNydDAJBgNVHRMEAjAAMIIBAwYKKwYBBAHWeQIEAgSB9ASB8QDvAHUAu9nfvB+KcbWTlCOXqpJ7RzhXlQqrUugakJZkNo4e0YUAAAFtBK0O6QAABAMARjBEAiAdmHDa5NbRtLx3lc9nQ9G81RZycaqQPMj3+sazAo5vjQIgLNuFD7zperowYJAtetRR4QUi/8dORH087fWBp+Waj5MAdgCHdb/nWXz4jEOZX73zbv9WjUdWNv9KtWDBtOr/XqCDDwAAAW0ErQ9SAAAEAwBHMEUCIQDzdkB41ukE5XQGDTp8N4r+Aw/TZ/FlhPrrZryVGz9RIQIgWiuG2RHKCbh6FtJo62ml9RDYHeW/xA7c5sBBeKkSfG4wCgYIKoZIzj0EAwIDSQAwRgIhALnmf8VUwhxU0dRo2iOlfRb9uFy3hXMceU4IEvsLSwOVAiEAxsfjpOn0JyE943lhWRvjXX8FOm927cI5mbZ5F+p6dAA=\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIID4zCCAsugAwIBAgIQBz/JpHsGAhj24Khq6fw+OzANBgkqhkiG9w0BAQsFADBhMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBDQTAeFw0xNzEyMDgxMjI4NTdaFw0yNzEyMDgxMjI4NTdaMHIxCzAJBgNVBAYTAkNOMSUwIwYDVQQKExxUcnVzdEFzaWEgVGVjaG5vbG9naWVzLCBJbmMuMR0wGwYDVQQLExREb21haW4gVmFsaWRhdGVkIFNTTDEdMBsGA1UEAxMUVHJ1c3RBc2lhIFRMUyBFQ0MgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASdQvDzv44jBee0APcvKOWszZsRjc4j+L6DLlYOf9tSgvfOJplfMeDNDZzOQEcJbVPD+yekJQUmObCPOrgMhqMIo4IBTzCCAUswHQYDVR0OBBYEFBKGRGYmCFQmj2U3silOJiHgk77bMB8GA1UdIwQYMBaAFAPeUDVW0Uy7ZvCj4hsbw5eyPdFVMA4GA1UdDwEB/wQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwEgYDVR0TAQH/BAgwBgEB/wIBADA0BggrBgEFBQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmRpZ2ljZXJ0LmNvbTBCBgNVHR8EOzA5MDegNaAzhjFodHRwOi8vY3JsMy5kaWdpY2VydC5jb20vRGlnaUNlcnRHbG9iYWxSb290Q0EuY3JsMEwGA1UdIARFMEMwNwYJYIZIAYb9bAECMCowKAYIKwYBBQUHAgEWHGh0dHBzOi8vd3d3LmRpZ2ljZXJ0LmNvbS9DUFMwCAYGZ4EMAQIBMA0GCSqGSIb3DQEBCwUAA4IBAQBZcGGhLE09CbQD5xP93NAuNC85G1BMa1OG2Q01TWvvgp7Qt1wNfRLAnhQT5pb7kRs+E7nM4IS894ufmuL452q8gYaq5HmvOmfhXMmL6K+eICfvyqjb/tSi8iy20ULO/TZhLhPor9tle52Yx811FG4i5vqwPIUEOEJ7pXe6RPVoBiwi4rbLspQGD/vYqrj9OJV4JctoIhhGq+y/sozU6nBXHfhVSD3x+hkOOst6tyRq481IyUWQHcFtwda3gfMnaA3dsag2dtJz33RIJIUfxXmVK7w4YzHOHifn7TYk8iNrDDLtql6vS8FjiUx3kJnI6zge1C9lUHhZ/aD3RiTJrwWI\n-----END CERTIFICATE-----"
  cert_private_data = "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgp8yx31T7g0TyZcU4IdJS4px8p0b9FOHqx0uIMwtIjP6gCgYIKoZIzj0DAQehRANCAAR+aGvOdizh+oAWwT6829WdcZw7oBJVU1UvKQdm7dW/7SIdrMEWq6NIWaERMKkLD6gQ6Y5KFV9oDQdSocGBtBvL\n-----END PRIVATE KEY-----"
}

resource "baiducloud_cert" "default" {
  cert_name         = "%s"
  cert_server_data  = baiducloud_cert.default2.cert_server_data
  cert_private_data = baiducloud_cert.default2.cert_private_data
}

data "baiducloud_certs" "default" {
  cert_name = baiducloud_cert.default.cert_name

  filter {
    name = "cert_name"
    values = ["test-BaiduAcc*"]
  }
}
`, BaiduCloudTestResourceAttrNamePrefix+"CertTest",
		BaiduCloudTestResourceAttrNamePrefix+"Cert")

}
