variable "unix_timestamp" {
  type   = string
}

# GCPプロジェクトの作成
resource "google_project" "bread" {
  provider   = google-beta.no_user_project_override
  # Organization > Folders > Project > Resource
  # See: https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy?hl=ja
  # folder_id  = "miyamo2theppl"
  name       = "Bread"
  project_id = "bread-${var.unix_timestamp}"

  # Firebaseのプロジェクトとして表示
  labels = {
    "firebase" = "enabled"
  }
}

# 依存APIの有効化
resource "google_project_service" "bread" {
  provider = google-beta.no_user_project_override
  project  = google_project.bread.project_id
  for_each = toset([
    "cloudresourcemanager.googleapis.com",
    "serviceusage.googleapis.com",
    "firestore.googleapis.com",
    "firebaserules.googleapis.com",
    "firebase.googleapis.com",
  ])
  service = each.key

  # Don't disable the service if the resource block is removed by accident.
  disable_on_destroy = false
}

# Firebaseプロジェクト
resource "google_firebase_project" "bread" {
  provider = google-beta.no_user_project_override
  project  = google_project.bread.project_id
}

# Firestore
resource "google_firestore_database" "bread" {
  provider                    = google-beta
  project                     = google_project.bread.project_id
  name                        = "(default)"
  # See available locations: https://firebase.google.com/docs/projects/locations#default-cloud-location
  location_id                 = "asia-northeast1"
  # "FIRESTORE_NATIVE" is required to use Firestore with Firebase SDKs, authentication, and Firebase Security Rules.
  type                        = "FIRESTORE_NATIVE"
  concurrency_mode            = "OPTIMISTIC"

  depends_on = [
    google_firebase_project.bread,
  ]
}

# Creates a ruleset of Firestore Security Rules from a local file.
resource "google_firebaserules_ruleset" "bread" {
  provider = google-beta
  project  = google_project.bread.project_id
  source {
    files {
      name = "firestore.rules"
      # Write security rules in a local file named "firestore.rules".
      # Learn more: https://firebase.google.com/docs/firestore/security/get-started
      content = file("firestore.rules")
    }
  }

  depends_on = [
    google_firestore_database.bread,
  ]
}

# Releases the ruleset for the Firestore instance.
resource "google_firebaserules_release" "bread" {
  provider     = google-beta
  name         = "cloud.firestore"  # must be cloud.firestore
  ruleset_name = google_firebaserules_ruleset.bread.name
  project      = google_project.bread.project_id

  depends_on = [
    google_firestore_database.bread,
  ]
}



