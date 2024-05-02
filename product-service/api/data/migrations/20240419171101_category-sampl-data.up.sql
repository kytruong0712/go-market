INSERT INTO public.categories(id, parent_id, category_name, category_code, description, is_navgitation, is_filtering, status)
VALUES (1, null, 'Điện thoại, Tablet', 'dien-thoai-tablet', 'Điện thoại, Tablet', true, true, 'ACTIVE'),
       (2, null, 'Laptop', 'laptop', 'Laptop', true, true, 'ACTIVE'),
       (3, null, 'Âm thanh', 'am-thanh', 'Âm thanh', true, true, 'ACTIVE'),
       (4, 1, 'Thương hiệu điện thoại', 'thuong-hieu-dien-thoai', 'Thương hiệu điện thoại', true, true, 'ACTIVE'),
       (5, 1, 'Thương hiệu Tablet', 'thuong-hieu-tablet', 'Thương hiệu Tablet', true, true, 'ACTIVE'),
       (6, 2, 'Chọn theo hãng', 'chon-theo-hang', 'Chọn theo hãng', true, true, 'ACTIVE'),
       (7, 2, 'Tầm giá', 'tam-gia', 'Tầm giá', true, true, 'ACTIVE'),
       (8, 3, 'Chọn loại tai nghe', 'chon-loai-tai-nghe', 'Chọn loại tai nghe', true, true, 'ACTIVE'),
       (9, 3, 'Hãng tai nghe', 'hang-tai-nghe', 'Hãng tai nghe', true, true, 'ACTIVE'),
       (10, 4, 'Apple', 'apple', 'Apple', true, true, 'ACTIVE'),
       (11, 4, 'Samsung', 'samsung', 'Samsung', true, true, 'ACTIVE'),
       (12, 5, 'Apple', 'apple', 'Apple', true, true, 'ACTIVE'),
       (13, 5, 'Samsung', 'samsung', 'Samsung', true, true, 'ACTIVE'),
       (14, 6, 'Dell', 'dell', 'Dell', true, true, 'ACTIVE'),
       (15, 6, 'Lenovo', 'lenovo', 'Lenovo', true, true, 'ACTIVE'),
       (16, 7, 'Dưới 10 triệu', 'duoi-10-trieu', 'Dưới 10 triệu', true, true, 'ACTIVE'),
       (17, 7, 'Từ 10 - 15 triệu', '10-15-trieu', 'Từ 10 - 15 triệu', true, true, 'ACTIVE'),
       (18, 8, 'Tai nghe true-wireless', 'tai-nghe-true-wireless', 'Tai nghe true-wireless', true, true, 'ACTIVE'),
       (19, 8, 'Tai nghe có dây', 'tai-nghe-co-day', 'Tai nghe có dây', true, true, 'ACTIVE'),
       (20, 9, 'Sennheiser', 'sennheiser', 'Sennheiser', true, true, 'ACTIVE'),
       (21, 9, 'Sony', 'sony', 'Sony', true, true, 'ACTIVE'),
       (22, 10, 'Iphone 15 series', 'iphone-15-series', 'Iphone 15 series', true, true, 'ACTIVE'),
       (23, 14, 'Dell XPS series', 'dell-xps-series', 'Dell XPS series', true, true, 'ACTIVE');

INSERT INTO public.uploaded_files(id, file_name, file_size_kb, file_path, file_type, resolution, status)
VALUES (1, 'cate1.png', '1000', 'foo/bar', 'IMAGE', '100x100', 'ACTIVE'),
       (2, 'cate2.png', '1000', 'foo/bar', 'IMAGE', '100x100', 'ACTIVE'),
       (3, 'cate3.png', '1000', 'foo/bar', 'IMAGE', '100x100', 'ACTIVE');

INSERT INTO public.category_images(id, category_id, uploaded_file_id, status, image_type)
VALUES (1, 1, 1, 'ACTIVE', 'CATEGORY_MENU_ICON'),
       (2, 2, 2, 'ACTIVE', 'CATEGORY_MENU_ICON'),
       (3, 3, 3, 'ACTIVE', 'CATEGORY_MENU_ICON');
