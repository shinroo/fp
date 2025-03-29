-- Collections
INSERT INTO "public"."directus_collections" ("collection", "icon", "note", "display_template", "hidden", "singleton", "translations", "archive_field", "archive_app_filter", "archive_value", "unarchive_value", "sort_field", "accountability", "color", "item_duplication_fields", "sort", "group", "collapse", "preview_url", "versioning") VALUES
('dog', NULL, NULL, NULL, 'f', 'f', NULL, NULL, 't', NULL, NULL, NULL, 'all', NULL, NULL, NULL, NULL, 'open', NULL, 'f') ON CONFLICT (collection) DO NOTHING;

-- Fields
INSERT INTO "public"."directus_fields" ("id", "collection", "field", "special", "interface", "options", "display", "display_options", "readonly", "hidden", "sort", "width", "translations", "note", "conditions", "required", "group", "validation", "validation_message") VALUES
(1, 'dog', 'identifier', NULL, 'input', '{"iconLeft":"grid_3x3"}', 'formatted-value', NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 't', NULL, NULL, NULL),
(2, 'dog', 'name', NULL, NULL, NULL, NULL, NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 'f', NULL, NULL, NULL),
(3, 'dog', 'gender', NULL, NULL, NULL, 'raw', NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 't', NULL, '{"_and":[{"_or":[{"gender":{"_eq":"Male"}},{"gender":{"_eq":"Female"}}]}]}', 'Gender must be either "Male" or "Female"'),
(4, 'dog', 'life_stage', NULL, NULL, NULL, NULL, NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 'f', NULL, '{"_and":[{"_or":[{"life_stage":{"_eq":"Puppy"}},{"life_stage":{"_eq":"Adult"}},{"life_stage":{"_eq":"Senior"}}]}]}', 'Life stage must be either "Puppy", "Adult" or "Senior"'),
(5, 'dog', 'image_url', NULL, NULL, NULL, 'formatted-value', NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 't', NULL, NULL, NULL),
(6, 'dog', 'breed', NULL, 'select-dropdown-m2o', '{"template":"{{name}}"}', NULL, NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 't', NULL, NULL, NULL),
(7, 'dog', 'spca_id', NULL, 'select-dropdown-m2o', '{"template":"{{name}}"}', NULL, NULL, 'f', 'f', NULL, 'full', NULL, NULL, NULL, 't', NULL, NULL, NULL),
(8, 'dog', 'created_at', NULL, 'datetime', NULL, NULL, NULL, 't', 'f', NULL, 'full', NULL, NULL, NULL, 't', NULL, NULL, NULL)
ON CONFLICT (id) DO NOTHING;

-- Presets
INSERT INTO "public"."directus_presets" ("id", "bookmark", "user", "role", "collection", "search", "layout", "layout_query", "layout_options", "refresh_interval", "filter", "icon", "color") VALUES
(1, NULL, 'd3b94bc3-36ad-432e-894e-33140614d412', NULL, 'dog', NULL, 'tabular', '{"tabular":{"fields":["identifier","gender","life_stage","name","image_url","spca_id"]}}', '{"cards":{"title":"{{name}}","subtitle":"{{identifier}}"},"tabular":{"widths":{}}}', NULL, NULL, 'bookmark', NULL) ON CONFLICT (id) DO NOTHING;

-- Settings
INSERT INTO "public"."directus_settings" ("id", "project_name", "project_url", "project_color", "project_logo", "public_foreground", "public_background", "public_note", "auth_login_attempts", "auth_password_policy", "storage_asset_transform", "storage_asset_presets", "custom_css", "storage_default_folder", "basemaps", "mapbox_key", "module_bar", "project_descriptor", "default_language", "custom_aspect_ratios", "public_favicon", "default_appearance", "default_theme_light", "theme_light_overrides", "default_theme_dark", "theme_dark_overrides", "report_error_url", "report_bug_url", "report_feature_url", "public_registration", "public_registration_verify_email", "public_registration_role", "public_registration_email_filter") VALUES
(1, 'Directus', NULL, '#233661', '82f28752-b7dd-4268-9dc0-8ed945fe3bb9', NULL, NULL, NULL, 25, NULL, 'all', NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'en-US', NULL, NULL, 'light', 'Directus Default', NULL, NULL, NULL, NULL, NULL, NULL, 'f', 't', NULL, NULL) ON CONFLICT (id) DO NOTHING;