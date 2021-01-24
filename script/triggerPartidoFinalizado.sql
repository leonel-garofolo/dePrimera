DROP TRIGGER addPointToTeam;

DELIMITER //
CREATE TRIGGER addPointToTeam AFTER UPDATE ON partidos
FOR EACH ROW
BEGIN
declare idLiga INT;
declare idCampeonato INT;
declare idPartidos INT;
declare idEquipoLocal INT;
declare resultadoLocal INT;
declare idEquipoVisitante INT;
declare resultadoVisitante INT;
   -- call debug_on('addPointToTeam');

   IF !(NEW.finalizado = OLD.finalizado) THEN
	  -- call debug_insert('addPointToTeam','change trigger');
      
      select 
		c.id_liga, 
        c.id_campeonato, 
        p.id_partidos, p.id_equipo_local, p.resultado_local, p.id_equipo_visitante, p.resultado_visitante
		from partidos p
		inner join campeonatos c on c.id_campeonato = p.id_campeonato
		where p.id_partidos = old.id_partidos 
        into idLiga, idCampeonato, idPartidos, idEquipoLocal, resultadoLocal, idEquipoVisitante, resultadoVisitante
        ;
      -- call debug_insert('addPointToTeam', concat('idLiga: ', CONVERT(idLiga,char), '| idCampeonato: ', CONVERT(idCampeonato,char), '| idPartidos: ', CONVERT(idPartidos,char), '| resultadoLocal: ', CONVERT(resultadoLocal,char), '| resultadoVisitante: ', CONVERT(resultadoVisitante,char)));  
      if(resultadoLocal = resultadoVisitante) then  
		update campeonatos_equipos set p_emp = p_emp + 1, puntos = puntos + 1 where id_liga = idLiga and id_campeonato = idCampeonato and id_equipo = idEquipoLocal;
        update campeonatos_equipos set p_emp = p_emp + 1, puntos = puntos + 1 where id_liga = idLiga and id_campeonato = idCampeonato and id_equipo = idEquipoVisitante;
      end if;  
	  if(resultadoLocal > resultadoVisitante) then  
		update campeonatos_equipos set p_gan = p_gan + 1, puntos = puntos + 3 where id_liga = idLiga and id_campeonato = idCampeonato and id_equipo = idEquipoLocal;
        update campeonatos_equipos set p_per = p_per + 1 where id_liga = idLiga and id_campeonato = idCampeonato and id_equipo = idEquipoVisitante;
      end if;       
      if(resultadoLocal < resultadoVisitante) then  
		update campeonatos_equipos set p_per = p_per + 1 where id_liga = idLiga and id_campeonato = idCampeonato and id_equipo = idEquipoLocal;
        update campeonatos_equipos set p_gan = p_gan + 1, puntos = puntos + 3 where id_liga = idLiga and id_campeonato = idCampeonato and id_equipo = idEquipoVisitante;
      end if;  
   END IF;
   -- call debug_off('addPointToTeam');
END;//
DELIMITER ;