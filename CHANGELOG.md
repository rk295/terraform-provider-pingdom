## Unreleased

## 1.1.4 (Mar 22, 2021)

* [`a57a29f`](https://github.com/rk295/terraform-provider-pingdom/commit/a57a29f) Update the fork with Nordcloud references (#1)
* [`653f94f`](https://github.com/rk295/terraform-provider-pingdom/commit/653f94f) Merge pull request russellcardullo#93 from marcelomilera/feature/add-pingdom-teams-data-source
* [`0b913da`](https://github.com/rk295/terraform-provider-pingdom/commit/0b913da) Merge pull request russellcardullo#92 from marcelomilera/bugfix/fix-pingdom-contact-data-source
* [`298beab`](https://github.com/rk295/terraform-provider-pingdom/commit/298beab) Merge pull request russellcardullo#86 from unprofession-al/fix_doc
* [`c815572`](https://github.com/rk295/terraform-provider-pingdom/commit/c815572) Merge pull request russellcardullo#89 from marcuz/patch-1
* [`c13cf2f`](https://github.com/rk295/terraform-provider-pingdom/commit/c13cf2f) Add pingdom_team data source
* [`59578a3`](https://github.com/rk295/terraform-provider-pingdom/commit/59578a3) Fix pingdom_contact data source returning wrong ID
* [`0b1a159`](https://github.com/rk295/terraform-provider-pingdom/commit/0b1a159) DOCS: fix SMS providers names
* [`4b59538`](https://github.com/rk295/terraform-provider-pingdom/commit/4b59538) README: fixed attribut userids for pingdom_check
* [`d90b4e2`](https://github.com/rk295/terraform-provider-pingdom/commit/d90b4e2) Merge pull request russellcardullo#81 from unprofession-al/fix_contact_creation
* [`971648d`](https://github.com/rk295/terraform-provider-pingdom/commit/971648d) Merge pull request russellcardullo#80 from zburgermeiszter/fix-readme
* [`5040d49`](https://github.com/rk295/terraform-provider-pingdom/commit/5040d49) Merge pull request russellcardullo#79 from unprofession-al/master
* [`b1b6487`](https://github.com/rk295/terraform-provider-pingdom/commit/b1b6487) contacts: set the high and low severity flags correctly in case of email notifications
* [`ff7c4b8`](https://github.com/rk295/terraform-provider-pingdom/commit/ff7c4b8) doc: fixed breaking change userids => member_ids
* [`4f88b13`](https://github.com/rk295/terraform-provider-pingdom/commit/4f88b13) README: Provider config fixed
* [`6de7cb4`](https://github.com/rk295/terraform-provider-pingdom/commit/6de7cb4) Update README.md
* [`35fc530`](https://github.com/rk295/terraform-provider-pingdom/commit/35fc530) doc: updated to state support for API 3.1

## 1.1.3 (October 20, 2020)

BREAKING CHANGES:

  * This release removes support for the deprecated v2.X Pingdom APIs.

NEW FEATURES:

  * Add support for the v3.1 Pingdom API (#77)

IMPROVEMENTS:
  
  * Documentation improvements (#73, #76)
  * Add GitHub actions workflows for linting and testing (#75)

## 1.1.2 (September 13, 2020)

NEW FEATURES:

  * Add support for managing teams, contacts, users (#36)
  * Allow adding users to teams (#61)

IMPROVEMENTS:

  * Add responsetimethreshold to checks (#36)
  * CI improvements (#48)
  * Documentation improvements (#45)
  * Uses latest patch version of go in CI builds (#50)
  * Update to terraform 0.12.18 (#51, #54)
  * Migrate to terraform-plugin-sdk (#65)
  * Sort tags on write to prevent unnecessary diffs (#58)
  * Documentation improvements (#53, #66, #67)
  * Use GitHub actions for builds and releases (#72)

BUG FIXES:
  * Include existing probefilter values on reads (#47)
  * Fix issue importing contacts (#60)

## 1.1.1 (October 5, 2019)

IMPROVEMENTS:

  * Add sensitive flags for secret values (#44)
  * Publish releases from Travis CI (#41)

## 1.0.0 (July 3, 2019)

NEW FEATURES:

  * Add TCP Checks (#21)
  * Add support for Public Reports (#21)
  * Add support for integrations/webhooks in checks (#14)
  * Add support for probe filters (#13)
  * Checks support paused parameter (#22)
  * Add support for tags on checks (#8)
  * Add support for importing existing checks (#9)
  * Add contact IDs to schema (#3)
  * Add option to use legacy notifications
  * Add spport for multi-user accounts (#1)

IMPROVEMENTS:

  * Use go modules to manage dependencies (#30)
  * Update to go-pingdom v1.0.0
  * Update to terraform 0.12.3 (#38)
  * Documentation updates (#12)

BUG FIXES:

  * Cannot use imported resource without forced re-creation (#31)
  * Fix teams response (#27)
  * Stop using `id` check schema in (#11)
  * Add default value for check URLs (#4)

## 0.2.0 (October 17, 2014)

IMPROVEMENTS:

  * Add support for Terraform 0.3.0

## 0.1.1 (September 16, 2014)

FEATURES:

  * Add support for managing ping type checks

BUG FIXES:

  * Don't override check resource values on create

## 0.1.0 (September 7, 2014)

  * Initial release
