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
select * from app_ppaises;
select * from app_provincias;
select * from comentarios;


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
    p.suspendido, p.iniciado, p.finalizado,   
    (select string_agg(aux_jug.nro_camiseta::text, ' ') 
		from campeonatos_goleadores aux_cg 
        inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_local.id_equipo
        where aux_cg.id_partido = p.id_partidos) as goleadores_local,
     (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from campeonatos_goleadores aux_cg 
        inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_visit.id_equipo
        where aux_cg.id_partido = p.id_partidos) as goleadores_visit ,
	(select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo
		where aux_sj.id_sancion= 2 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_amarillas,
    (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo
		where aux_sj.id_sancion= 1 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_rojas,
    (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo
		where aux_sj.id_sancion= 2 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_amarillas,
    (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo
		where aux_sj.id_sancion= 1 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_rojas    
from partidos p
inner join ligas l on l.id_liga = p.id_liga
inner join campeonatos c on c.id_campeonato = p.id_campeonato
inner join equipos e_local on e_local.id_equipo = p.id_equipo_local
inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante
left join arbitros a on a.id_arbitro = p.id_arbitro
left join asistentes asis on asis.id_asistente = p.id_asistente
where fecha_encuentro = '2020-08-07%';

-- query for get all partidos from Equipo.
select p.id_partidos, p.fecha_encuentro,
	l.nombre as ligaName, c.descripcion as campeonatoName, 
    e_local.nombre as eLocalName, e_visit.nombre as eVisitName, 
    p.resultado_local, p.resultado_visitante,
    p.suspendido, p.iniciado, p.finalizado
from partidos p
inner join ligas l on l.id_liga = p.id_liga
inner join campeonatos c on c.id_campeonato = p.id_campeonato
inner join equipos e_local on e_local.id_equipo = p.id_equipo_local
inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante
left join arbitros a on a.id_arbitro = p.id_arbitro
left join asistentes asis on asis.id_asistente = p.id_asistente
where e_local.id_equipo = 2 or e_visit.id_equipo = 2;

select p.id_partidos, p.fecha_encuentro,
	l.nombre as ligaName, c.descripcion as campeonatoName, 
    e_local.nombre as eLocalName, e_visit.nombre as eVisitName, 
    p.resultado_local, p.resultado_visitante,
    p.suspendido, p.iniciado, p.finalizado,   
    (select array_to_string(array(aux_jug.nro_camiseta), ' ')
		from campeonatos_goleadores aux_cg 
        inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_local.id_equipo
        where aux_cg.id_partido = p.id_partidos) as goleadores_local    
from partidos p
inner join ligas l on l.id_liga = p.id_liga
inner join campeonatos c on c.id_campeonato = p.id_campeonato
inner join equipos e_local on e_local.id_equipo = p.id_equipo_local
inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante
left join arbitros a on a.id_arbitro = p.id_arbitro
left join asistentes asis on asis.id_asistente = p.id_asistente
where c.id_campeonato = 1
order by fecha_encuentro asc;


select string_agg(aux_jug.nro_camiseta::text, ',') 
		from campeonatos_goleadores aux_cg 
        inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = 1
        where aux_cg.id_partido = 1;

-- get FIXTURE
select p.id_partidos, p.fecha_encuentro,
	l.nombre as ligaName, c.descripcion as campeonatoName, 
    e_local.nombre as eLocalName, e_visit.nombre as eVisitName, 
    p.resultado_local, p.resultado_visitante,
    p.suspendido, p.iniciado, p.finalizado,   
    (select string_agg(aux_jug.nro_camiseta::text, ' ') 
		from campeonatos_goleadores aux_cg 
        inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_local.id_equipo
        where aux_cg.id_partido = p.id_partidos) as goleadores_local,
     (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from campeonatos_goleadores aux_cg 
        inner join jugadores aux_jug on aux_jug.id_jugadores = aux_cg.id_jugadores and aux_jug.id_equipo = e_visit.id_equipo
        where aux_cg.id_partido = p.id_partidos) as goleadores_visit ,
	(select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo
		where aux_sj.id_sancion= 2 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_amarillas,
    (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_local.id_equipo
		where aux_sj.id_sancion= 1 and aux_sj.id_partidos = p.id_partidos) as sanciones_local_rojas,
    (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo
		where aux_sj.id_sancion= 2 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_amarillas,
    (select string_agg(aux_jug.nro_camiseta::text, ' ')
		from sanciones_jugadores aux_sj 
		inner join jugadores aux_jug on aux_jug.id_jugadores = aux_sj.id_jugador and aux_jug.id_equipo = e_visit.id_equipo
		where aux_sj.id_sancion= 1 and aux_sj.id_partidos = p.id_partidos) as sanciones_visit_rojas    
from partidos p
inner join ligas l on l.id_liga = p.id_liga
inner join campeonatos c on c.id_campeonato = p.id_campeonato
inner join equipos e_local on e_local.id_equipo = p.id_equipo_local
inner join equipos e_visit on e_visit.id_equipo = p.id_equipo_visitante
left join arbitros a on a.id_arbitro = p.id_arbitro
left join asistentes asis on asis.id_asistente = p.id_asistente
where c.id_campeonato = 1
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
select p.nombre as p_nombre, p.apellido as p_apellido, e.nombre as e_nombre, 
	(case when sj.id_sancion = 1 then count(sj.id_sancion) else 0 end ) as c_rojas,
    (case when sj.id_sancion = 2 then count(sj.id_sancion) else 0 end ) as c_amarillas,
    (case when sj.id_sancion = 3 then count(sj.id_sancion) else 0 end ) as c_azules
from sanciones_jugadores sj
inner join partidos partido on partido.id_partidos = sj.id_partidos
inner join jugadores j on j.id_jugadores = sj.id_jugador
inner join equipos e on e.id_equipo = j.id_equipo
inner join personas p on p.id_persona = j.id_persona
where partido.id_campeonato = 1
group by p.nombre, p.apellido, e.nombre, sj.id_sancion
order by p.apellido asc, p.nombre asc ;

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
    
    
-- get query of notification
select 
	(select count(*) from ligas) as ligas,
    (select count(*) from campeonatos) as campeonatos,
    (select count(*) from equipos) as equipos,
    (select count(*) from arbitros) as arbitros,
    (select count(*) from asistentes) as asistentes,
    (select count(*) from jugadores) as jugadores
    ;
 


-- get campeonatos from asistentes
select c.*
from campeonatos c
inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato
inner join asistentes a on a.id_campeonato = ce.id_campeonato
inner join personas p on p.id_persona = a.id_persona
where p.id_user= 'asis1' and p.idgrupo = 2;

-- get campeonatos from jugadores
select c.*
from campeonatos c
inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato
inner join jugadores j on j.id_equipo = ce.id_equipo
inner join personas p on p.id_persona = j.id_persona
inner join app_users au on au.id_user = p.id_user
where au.id_user= 'jug1' and p.idgrupo = 3;
 
 -- get campeonatos from arbitros
select c.*
from campeonatos c
inner join campeonatos_equipos ce on ce.id_campeonato = c.id_campeonato
inner join arbitros a on a.id_campeonato = ce.id_campeonato
inner join personas p on p.id_persona = a.id_persona
inner join app_users au on au.id_user = p.id_user
where au.id_user= 'arb1' and p.idgrupo = 4;

-- get grupo from user
select p.idgrupo, ag.descripcion 
from personas p
inner join app_grupos ag on p.idgrupo = ag.idgrupo
where p.id_user = 'jug1';

select * from equipos;

-- get equipos from asistentes
select e.id_equipo, e.nombre, e.habilitado, e.foto 
from equipos e
inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo
inner join asistentes a on a.id_campeonato = ce.id_campeonato
inner join personas p on p.id_persona = a.id_persona
where p.id_user= 'asis1' and p.idgrupo = 2;

-- get equipos from jugadores
select e.*
from equipos e
inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo
inner join jugadores j on j.id_equipo = ce.id_equipo
inner join personas p on p.id_persona = j.id_persona
inner join app_users au on au.id_user = p.id_user
where au.id_user= 'jug1' and p.idgrupo = 3;
 
 -- get equipos from arbitros
select e.*
from equipos e
inner join campeonatos_equipos ce on ce.id_equipo = e.id_equipo
inner join arbitros a on a.id_campeonato = ce.id_campeonato
inner join personas p on p.id_persona = a.id_persona
where p.id_user= 'arb1' and p.idgrupo = 4;

-- get Plantel from equipo
select j.id_jugadores, p.apellido, p.nombre, j.nro_camiseta 
from jugadores j
inner join personas p on j.id_persona = p.id_persona
where p.apellido is not null and p.nombre is not null and j.id_equipo = 2
order by p.apellido asc, p.nombre asc;

-- get jugadores of the equipos local and visit
select p.id_partidos, 
	jlocal.id_jugadores as jug_local, jlocal.nro_camiseta as nro_camiseta_local, 
    jvisit.id_jugadores as jug_visit, jvisit.nro_camiseta as nro_camiseta_visit
from partidos p
inner join jugadores jlocal on jlocal.id_equipo = p.id_equipo_local
inner join jugadores jvisit on jvisit.id_equipo = p.id_equipo_visitante
where id_partidos = 964;

-- get fecha de partidos proximos.
select fecha_encuentro from partidos p where p.fecha_encuentro > current_date group by fecha_encuentro  order by fecha_encuentro asc;


select c.id_liga,  c.id_campeonato, p.id_partidos, p.id_equipo_local, p.resultado_local, p.id_equipo_visitante, p.resultado_visitante
from partidos p
inner join campeonatos c on c.id_campeonato = p.id_campeonato
where p.id_partidos = 965 ;

select id_partidos, resultado_local, resultado_visitante, finalizado 
from partidos 
where id_partidos = 966
order by id_partidos asc;

update partidos set finalizado = 1 where id_partidos = 966;
