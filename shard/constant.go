package shard

const (
	SQLFUNCS = `
CREATE SEQUENCE ?shard.id_seq;

-- _next_id returns unique sortable id.
CREATE FUNCTION ?shard._next_id(tm timestamptz, shard_id int, seq_id bigint)
RETURNS bigint AS $$
DECLARE
  max_shard_id CONSTANT bigint := 2048;
  max_seq_id CONSTANT bigint := 4096;
  id bigint;
BEGIN
  shard_id := shard_id % max_shard_id;
  seq_id := seq_id % max_seq_id;
  id := (floor(extract(epoch FROM tm) * 1000)::bigint - ?epoch) << 23;
  id := id | (shard_id << 12);
  id := id | seq_id;
  RETURN id;
END;
$$
LANGUAGE plpgsql
IMMUTABLE;

CREATE FUNCTION ?shard.next_id()
RETURNS bigint AS $$
BEGIN
   RETURN ?shard._next_id(clock_timestamp(), ?shard_id, nextval('?shard.id_seq'));
END;
$$
LANGUAGE plpgsql;
`
	CREATETABLETKB = `CREATE TABLE ?shard.tkbs (id bigint DEFAULT ?shard.next_id(), ma_mon_hoc text,ten_mon text, lop text, 
	khoa_nganh text, nganh text, nhom text, to_hop text, to_th text, thu text, kip text, so_cho_con_lai bigint, sy_so text, 
	phong text, nha text,hinh_thuc_thi text, ma_gv text, ten_gv text, ghi_chu text, ngay_bd timestamp with time zone, 
	ngay_kt timestamp with time zone, khoa text, bo_mon text, so_tc text, ts_tiet text, lt text, bt text, btl text, thtn text, tu_hoc text)`
	CREATETABLEREGISTSUBJECT = `CREATE TABLE ?shard.register_subject (id bigint DEFAULT ?shard.next_id(), ma_sv text, id_mon bigint, ma_mon_hoc text)`
)
