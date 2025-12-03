-- Допустим, у нас есть подготовленный statement
PREPARE user_plan AS SELECT * FROM users WHERE department = $1;

-- Приложение использует его неделями, всё быстро
EXECUTE user_plan('engineering');
EXECUTE user_plan('marketing');

-- Админ делает maintenance
VACUUM
ANALYZE users;
-- Или перезагружает конфиг
SELECT pg_reload_conf();
-- Или просто накатывает патч

-- Кэш планов сброшен!
-- Но приложение продолжает работать, не подозревая о проблеме
EXECUTE user_plan('sales');
-- ← Построение плана с нуля!

