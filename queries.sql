SET SQL_SAFE_UPDATES = 0;

select * from personas;
select * from arbitros;
select * from ligas;
select * from equipos;
select * from campeonatos;
select * from arbitros;
select * from asistentes;
select * from partidos;
select * from jugadores;
select * from campeonatos;

-- authentication
select * from app_users;

select * from ligas where id_liga = 1;

select id_campeonato, id_liga, id_modelo, descripcion, fecha_inicio, fecha_fin from campeonatos;

delete from arbitros where id_arbitro = 1 and id_persona = 1;

-- query for get all partidos from Date.
select p.id_partidos, p.fecha_encuentro,
	l.nombre as ligaName, c.descripcion as campeonatoName, 
    e_local.nombre as eLocalName, e_visit.nombre as eVisitName, 
    p.resultado_local, p.resultado_visitante,
    p.suspendido
from partidos p
inner join ligas l on l.id_liga = p.id_liga
inner join campeonatos c on c.id_campeonato = p.id_campeonato
inner join equipos e_local on e_local.id_equipo = p.id_equipo_local
inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante
inner join arbitros a on a.id_arbitro = p.id_arbitro
inner join asistentes asis on asis.id_asistente = p.id_asistente
where fecha_encuentro like "2020-07-04%";

-- query for get history of team.
select p.id_partidos, p.fecha_encuentro,
	l.nombre as ligaName, c.descripcion as campeonatoName, 
    e_local.nombre as eLocalName, e_visit.nombre as eVisitName, 
    p.resultado_local, p.resultado_visitante,
    p.suspendido
from partidos p
inner join ligas l on l.id_liga = p.id_liga
inner join campeonatos c on c.id_campeonato = p.id_campeonato
inner join equipos e_local on e_local.id_equipo = p.id_equipo_local
inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante
inner join arbitros a on a.id_arbitro = p.id_arbitro
inner join asistentes asis on asis.id_asistente = p.id_asistente
where id_equipo_local = 2 or id_equipo_visitante = 2;

delete from partidos;
select * from partidos;
-- insert partidos fixture
insert into partidos(id_liga, id_campeonato, id_equipo_local, id_equipo_visitante, fecha_encuentro ) 
values (
	1,
	1, 
    (select id_equipo from equipos where id_campeonato = 1 and nro_equipo = 1), 
    (select id_equipo from equipos where id_campeonato = 1 and nro_equipo = 2),
    current_timestamp()
    );