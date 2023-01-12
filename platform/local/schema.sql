DROP TABLE IF EXISTS shop.product;

CREATE TABLE shop.product (
    id varchar(36) NOT NULL PRIMARY KEY,
    available boolean DEFAULT true,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO `product` (`id`, `available`, `created_at`, `updated_at`)
VALUES
    ('22e04f8a-c18d-4f80-8a34-ebd26122274b', 1 , '2023-01-02 15:20:44', '2023-01-02 15:20:44'),
    ('fa7617c3-7247-4cc9-9047-c8111440728a', 1 , '2023-01-02 15:20:44', '2023-01-02 15:20:44'),
    ('7bd3c403-fd16-47fa-ba77-87412dcef1b0', 1 , '2023-01-02 15:20:44', '2023-01-02 15:20:44');
