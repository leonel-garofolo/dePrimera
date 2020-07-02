select * from personas;
select * from arbitros;
select * from ligas;
select * from equipos;
select * from campeonatos;
select * from arbitros;
select * from asistentes;
select * from partidos;

select * from ligas where id_liga = 1;

select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin from campeonatos;

delete from arbitros where id_arbitro = 1 and id_persona = 1;