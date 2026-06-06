-- =============================================================================
-- mall-admin 遗留 GVA 数据清理脚本
-- 用途：删除已下线功能的物理表、菜单/API/Casbin 元数据，并补全商城模块种子
-- 使用：mysql -uroot -p gva < scripts/cleanup_legacy.sql
-- 建议：执行前备份数据库；执行后重启后端服务
-- =============================================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- -----------------------------------------------------------------------------
-- 1. 清理角色-菜单关联（先于菜单删除）
-- -----------------------------------------------------------------------------
DELETE sam
FROM sys_authority_menus sam
INNER JOIN sys_base_menus m ON sam.sys_base_menu_id = m.id
WHERE m.name IN ('example', 'upload', 'breakpoint', 'customer', 'formCreate', 'anInfo');

-- -----------------------------------------------------------------------------
-- 2. 清理菜单按钮与参数
-- -----------------------------------------------------------------------------
DELETE smb
FROM sys_base_menu_btns smb
INNER JOIN sys_base_menus m ON smb.sys_base_menu_id = m.id
WHERE m.name IN ('example', 'upload', 'breakpoint', 'customer', 'formCreate', 'anInfo');

DELETE smp
FROM sys_base_menu_parameters smp
INNER JOIN sys_base_menus m ON smp.sys_base_menu_id = m.id
WHERE m.name IN ('example', 'upload', 'breakpoint', 'customer', 'formCreate', 'anInfo');

-- -----------------------------------------------------------------------------
-- 3. 删除遗留菜单（先子后父）
-- -----------------------------------------------------------------------------
DELETE FROM sys_base_menus
WHERE name IN ('upload', 'breakpoint', 'customer', 'formCreate', 'anInfo');

DELETE FROM sys_base_menus WHERE name = 'example';

-- -----------------------------------------------------------------------------
-- 4. 删除遗留 API 登记
-- -----------------------------------------------------------------------------
DELETE FROM sys_apis
WHERE path LIKE '/customer/%'
   OR path LIKE '/info/%'
   OR path IN (
        '/fileUploadAndDownload/findFile',
        '/fileUploadAndDownload/breakpointContinue',
        '/fileUploadAndDownload/breakpointContinueFinish',
        '/fileUploadAndDownload/removeChunk'
   )
   OR api_group IN ('客户', '公告', '分片上传');

-- -----------------------------------------------------------------------------
-- 5. 删除 Casbin 策略
-- -----------------------------------------------------------------------------
DELETE FROM casbin_rule
WHERE v1 LIKE '/customer/%'
   OR v1 LIKE '/info/%'
   OR v1 IN (
        '/fileUploadAndDownload/findFile',
        '/fileUploadAndDownload/breakpointContinue',
        '/fileUploadAndDownload/breakpointContinueFinish',
        '/fileUploadAndDownload/removeChunk'
   );

-- -----------------------------------------------------------------------------
-- 6. 清理公开 API 白名单中的公告接口，补全当前需要的公开路由
-- -----------------------------------------------------------------------------
DELETE FROM sys_ignore_apis
WHERE path LIKE '/info/%'
   OR path LIKE '/autoCode/%'
   OR path = '/api/freshCasbin';

INSERT INTO sys_ignore_apis (created_at, updated_at, method, path)
SELECT NOW(3), NOW(3), t.method, t.path
FROM (
    SELECT 'POST' AS method, '/site/auth/register' AS path UNION ALL
    SELECT 'POST', '/site/auth/login' UNION ALL
    SELECT 'POST', '/site/auth/logout' UNION ALL
    SELECT 'GET',  '/site/auth/profile' UNION ALL
    SELECT 'GET',  '/base/uploadConfig' UNION ALL
    SELECT 'GET',  '/health'
) AS t
WHERE NOT EXISTS (
    SELECT 1 FROM sys_ignore_apis s WHERE s.method = t.method AND s.path = t.path
);

-- -----------------------------------------------------------------------------
-- 7. 清理引用已下线接口的操作历史（可选，保持库干净）
-- -----------------------------------------------------------------------------
DELETE FROM sys_operation_records
WHERE path LIKE '/customer/%'
   OR path LIKE '/info/%'
   OR path IN (
        '/fileUploadAndDownload/findFile',
        '/fileUploadAndDownload/breakpointContinue',
        '/fileUploadAndDownload/breakpointContinueFinish',
        '/fileUploadAndDownload/removeChunk'
   );

-- -----------------------------------------------------------------------------
-- 8. 补全商城菜单（旧库可能缺失）
-- -----------------------------------------------------------------------------
INSERT INTO sys_base_menus (
    created_at, updated_at, menu_level, parent_id, path, name, hidden, component,
    sort, active_name, keep_alive, default_menu, title, icon, close_tab, transition_type
)
SELECT NOW(3), NOW(3), 0, 0, 'mall', 'mall', 0, 'view/mall/index.vue',
       2, '', 0, 0, '商城管理', 'shopping-cart', 0, ''
WHERE NOT EXISTS (SELECT 1 FROM sys_base_menus WHERE name = 'mall');

INSERT INTO sys_base_menus (
    created_at, updated_at, menu_level, parent_id, path, name, hidden, component,
    sort, active_name, keep_alive, default_menu, title, icon, close_tab, transition_type
)
SELECT NOW(3), NOW(3), 1, p.id, 'member', 'mallMember', 0, 'view/mall/member/index.vue',
       1, '', 0, 0, '会员管理', 'user', 0, ''
FROM sys_base_menus p
WHERE p.name = 'mall'
  AND NOT EXISTS (SELECT 1 FROM sys_base_menus WHERE name = 'mallMember');

INSERT IGNORE INTO sys_authority_menus (sys_base_menu_id, sys_authority_authority_id)
SELECT m.id, a.authority_id
FROM sys_base_menus m
CROSS JOIN (SELECT 888 AS authority_id UNION ALL SELECT 8881 UNION ALL SELECT 9528) a
WHERE m.name IN ('mall', 'mallMember');

-- -----------------------------------------------------------------------------
-- 9. 补全商城 API 与 Casbin
-- -----------------------------------------------------------------------------
INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
SELECT NOW(3), NOW(3), t.path, t.description, t.api_group, t.method
FROM (
    SELECT '/mall/member/list' AS path, '会员列表' AS description, '商城会员' AS api_group, 'POST' AS method UNION ALL
    SELECT '/mall/member/detail', '会员详情', '商城会员', 'POST' UNION ALL
    SELECT '/api/freshCasbin', '刷新casbin权限缓存', 'api', 'GET'
) AS t
WHERE NOT EXISTS (
    SELECT 1 FROM sys_apis s WHERE s.path = t.path AND s.method = t.method
);

INSERT IGNORE INTO casbin_rule (ptype, v0, v1, v2)
SELECT 'p', a.role, t.path, t.method
FROM (
    SELECT '/mall/member/list' AS path, 'POST' AS method UNION ALL
    SELECT '/mall/member/detail', 'POST' UNION ALL
    SELECT '/api/freshCasbin', 'GET'
) AS t
CROSS JOIN (SELECT '888' AS role UNION ALL SELECT '8881' UNION ALL SELECT '9528') AS a
WHERE NOT EXISTS (
    SELECT 1 FROM casbin_rule c
    WHERE c.ptype = 'p' AND c.v0 = a.role AND c.v1 = t.path AND c.v2 = t.method
);

-- -----------------------------------------------------------------------------
-- 10. 删除废弃物理表（保留上传与媒体库分类表）
-- -----------------------------------------------------------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
DROP TABLE IF EXISTS `exa_files`;
DROP TABLE IF EXISTS `exa_customers`;
DROP TABLE IF EXISTS `gva_announcements_info`;

-- -----------------------------------------------------------------------------
-- 11. 对齐菜单文案与图标（旧库残留）
-- -----------------------------------------------------------------------------
UPDATE sys_base_menus SET title = '系统工具' WHERE name = 'systemTools' AND title = '编程辅助';
UPDATE sys_base_menus SET icon = 'odometer' WHERE name = 'dashboard' AND icon = 'customer-gva';

SET FOREIGN_KEY_CHECKS = 1;

-- 完成。若需默认会员等级，请启动一次后端（AutoMigrate + seedMallDefaults）。
