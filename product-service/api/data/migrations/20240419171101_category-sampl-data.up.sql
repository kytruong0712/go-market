INSERT INTO public.categories(id, parent_id, category_name, category_code, description, is_navgitation, status, created_at)
VALUES (1, null, 'Điện thoại, Tablet', 'dien-thoai-tablet', 'Điện thoại, Tablet', true, 'ACTIVE', '2024-05-01 00:00:00.000000 +00:00'),
       (2, null, 'Laptop', 'laptop', 'Laptop', true, 'ACTIVE', '2024-05-02 00:00:00.000000 +00:00'),
       (3, null, 'Âm thanh', 'am-thanh', 'Âm thanh', true, 'ACTIVE', '2024-05-03 00:00:00.000000 +00:00'),
       (4, 1, 'Thương hiệu điện thoại', 'thuong-hieu-dien-thoai', 'Thương hiệu điện thoại', true, 'ACTIVE', '2024-05-04 00:00:00.000000 +00:00'),
       (5, 1, 'Thương hiệu Tablet', 'thuong-hieu-tablet', 'Thương hiệu Tablet', true, 'ACTIVE', '2024-05-05 00:00:00.000000 +00:00'),
       (6, 2, 'Chọn theo hãng', 'chon-theo-hang', 'Chọn theo hãng', true, 'ACTIVE', '2024-05-06 00:00:00.000000 +00:00'),
       (7, 2, 'Tầm giá', 'tam-gia', 'Tầm giá', true, 'ACTIVE', '2024-05-07 00:00:00.000000 +00:00'),
       (8, 3, 'Chọn loại tai nghe', 'chon-loai-tai-nghe', 'Chọn loại tai nghe', true, 'ACTIVE', '2024-05-08 00:00:00.000000 +00:00'),
       (9, 3, 'Hãng tai nghe', 'hang-tai-nghe', 'Hãng tai nghe', true, 'ACTIVE', '2024-05-09 00:00:00.000000 +00:00'),
       (10, 4, 'Apple', 'apple', 'Apple', true, 'ACTIVE', '2024-05-10 00:00:00.000000 +00:00'),
       (11, 4, 'Samsung', 'samsung', 'Samsung', true, 'ACTIVE', '2024-05-11 00:00:00.000000 +00:00'),
       (12, 5, 'Apple', 'apple', 'Apple', true, 'ACTIVE', '2024-05-12 00:00:00.000000 +00:00'),
       (13, 5, 'Samsung', 'samsung', 'Samsung', true, 'ACTIVE', '2024-05-13 00:00:00.000000 +00:00'),
       (14, 6, 'Dell', 'dell', 'Dell', true, 'ACTIVE', '2024-05-14 00:00:00.000000 +00:00'),
       (15, 6, 'Lenovo', 'lenovo', 'Lenovo', true, 'ACTIVE', '2024-05-15 00:00:00.000000 +00:00'),
       (16, 7, 'Dưới 10 triệu', 'duoi-10-trieu', 'Dưới 10 triệu', true, 'ACTIVE', '2024-05-16 00:00:00.000000 +00:00'),
       (17, 7, 'Từ 10 - 15 triệu', '10-15-trieu', 'Từ 10 - 15 triệu', true, 'ACTIVE', '2024-05-17 00:00:00.000000 +00:00'),
       (18, 8, 'Tai nghe true-wireless', 'tai-nghe-true-wireless', 'Tai nghe true-wireless', true, 'ACTIVE', '2024-05-18 00:00:00.000000 +00:00'),
       (19, 8, 'Tai nghe có dây', 'tai-nghe-co-day', 'Tai nghe có dây', true, 'ACTIVE', '2024-05-19 00:00:00.000000 +00:00'),
       (20, 9, 'Sennheiser', 'sennheiser', 'Sennheiser', true, 'ACTIVE', '2024-05-20 00:00:00.000000 +00:00'),
       (21, 9, 'Sony', 'sony', 'Sony', true, 'ACTIVE', '2024-05-21 00:00:00.000000 +00:00'),
       (22, 10, 'Iphone 15 series', 'iphone-15-series', 'Iphone 15 series', true, 'ACTIVE', '2024-05-22 00:00:00.000000 +00:00'),
       (23, 14, 'Dell XPS series', 'dell-xps-series', 'Dell XPS series', true, 'ACTIVE', '2024-05-23 00:00:00.000000 +00:00'),
       -- data for negative cases
       (24, null, 'Inactive Category', 'inactive-category', 'Inactive Category', true, 'INACTIVE', '2024-05-24 00:00:00.000000 +00:00'),
       (25, 24, 'Inactive Sub Category', 'inactive-sub-category', 'Inactive Sub Category', true, 'INACTIVE', '2024-05-25 00:00:00.000000 +00:00'),
       (26, null, 'Promotion event', 'promotion-event', 'Promotion event', false, 'ACTIVE', '2024-05-26 00:00:00.000000 +00:00'),
       (27, 4, 'Nokia', 'nokia', 'Nokia', false, 'ACTIVE', '2024-05-27 00:00:00.000000 +00:00');

INSERT INTO public.uploaded_files(id, file_name, file_size_kb, file_path, file_type, resolution, status)
VALUES (1, 'cate1.png', '1000', 'foo/bar', 'IMAGE', '100x100', 'ACTIVE'),
       (2, 'cate2.png', '1000', 'foo/bar', 'IMAGE', '100x100', 'ACTIVE'),
       (3, 'cate3.png', '1000', 'foo/bar', 'IMAGE', '100x100', 'ACTIVE');

INSERT INTO public.category_images(id, category_id, uploaded_file_id, status, image_type)
VALUES (1, 1, 1, 'ACTIVE', 'CATEGORY_MENU_ICON'),
       (2, 2, 2, 'ACTIVE', 'CATEGORY_MENU_ICON'),
       (3, 3, 3, 'ACTIVE', 'CATEGORY_MENU_ICON');
