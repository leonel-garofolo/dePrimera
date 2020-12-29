SET SQL_SAFE_UPDATES = 0;
drop table sanciones_jugadores;
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
select * from campeonatos_equipos;

insert into campeonatos_equipos(id_liga, id_campeonato, id_equipo)
select 2,2, id_equipo from equipos where id_equipo > 5;

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
left join arbitros a on a.id_arbitro = p.id_arbitro
left join asistentes asis on asis.id_asistente = p.id_asistente
where fecha_encuentro like "2020-12-26%";

-- get FIXTURE
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
left join arbitros a on a.id_arbitro = p.id_arbitro
left join asistentes asis on asis.id_asistente = p.id_asistente
where c.id_campeonato = 2
order by fecha_encuentro asc;

-- get Table Position
select e.nombre, 
	ce.nro_equipo,
	ce.puntos, ce.p_gan, ce.p_emp, ce.p_per
from campeonatos_equipos ce
inner join campeonatos c on c.id_campeonato = ce.id_campeonato
inner join equipos e on e.id_equipo = ce.id_equipo
where c.id_campeonato = 2
order by ce.puntos desc;

-- get Sanciones por Campeonatos
select p.apellido_nombre, e.nombre as e_nombre, 
	(case when sj.id_sancion = 1 then count(sj.id_sancion) else 0 end ) as c_rojas,
    (case when sj.id_sancion = 2 then count(sj.id_sancion) else 0 end ) as c_amarillas,
    (case when sj.id_sancion = 3 then count(sj.id_sancion) else 0 end ) as c_azules
from sanciones_jugadores sj
inner join jugadores j on j.id_jugadores = sj.id_jugador
inner join equipos e on e.id_equipo = j.id_equipo
inner join personas p on p.id_persona = j.id_persona
where sj.id_campeonato = 1
group by p.apellido_nombre, e.nombre, sj.id_sancion
order by p.apellido_nombre asc;

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