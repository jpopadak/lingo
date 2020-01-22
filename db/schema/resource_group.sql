
CREATE TABLE resource_group (
                                uuid                     BINARY(16)   NOT NULL,
                                uuid_text                VARCHAR(36)           GENERATED ALWAYS AS
                                                             (insert(
                                            insert(
                                                    insert(
                                                            insert(HEX(uuid), 9, 0, '-'),
                                                            14, 0, '-'),
                                                    19, 0, '-'),
                                            24, 0, '-')
                                                             ) VIRTUAL,
                                resource_group_type_uuid BINARY(16)   NOT NULL,
                                human_label              VARCHAR(128) NOT NULL,
                                description              VARCHAR(256) NOT NULL,
                                created_at               DATETIME(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
                                updated_at               DATETIME(6)  NOT NULL,
                                deleted_at               DATETIME(6),

                                PRIMARY KEY (uuid),
                                CONSTRAINT FOREIGN KEY fk_resource_group_type_uuid (resource_group_type_uuid) REFERENCES resource_group_type (uuid),
                                INDEX idx_created_at (created_at),
                                INDEX idx_updated_at (updated_at)
) DEFAULT CHARACTER SET = 'utf8mb4', DEFAULT COLLATE = 'utf8mb4_unicode_ci';
