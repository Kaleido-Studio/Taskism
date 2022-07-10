import Fastify from 'fastify';
import ILYFamily from 'types/ILYFamily';

const family: ILYFamily = { yhw: 'male', lx: 'female' };
const fastify = Fastify({
  logger: {
    transport: process.env.DEV
      ? {
          target: 'pino-pretty',
          options: {
            translateTime: 'HH:MM:ss Z',
            ignore: 'pid,hostname',
          },
        }
      : undefined,
  },
});

fastify.get('/', (req, res) => {
  res.send(family);
});

fastify.listen({ port: 3001 }, (err, addr) => {
  if (err) throw err;
  console.log(`server listening on ${addr}`);
});
