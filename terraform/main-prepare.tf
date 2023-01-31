## ------------------ Build packages and generate random uuids for func versions ------------

resource "null_resource" "build_zip" {
  triggers = {
    always_run = timestamp()
  }
  # Generate build.zip file
  provisioner "local-exec" {
    command = "cd ../; bash scripts/build-zip.sh"
  }
}

data "archive_file" "init" {
  depends_on = [
    null_resource.build_zip,
  ]

  type        = "zip"
  source_dir  = "../build/"
  output_path = "build.zip"
}
