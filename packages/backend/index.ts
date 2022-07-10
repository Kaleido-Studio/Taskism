import Fastify from 'fastify';
import ILYFamily, { Family } from 'types/ILYFamily';
import { TypeBoxTypeProvider } from '@fastify/type-provider-typebox';
import cors from '@fastify/cors';
import ETag from '@fastify/etag';

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
}).withTypeProvider<TypeBoxTypeProvider>();
fastify.register(cors);
fastify.register(ETag);

fastify.get<{ Reply: ILYFamily }>('/', { schema: { response: { 200: Family } } }, (req, res) => {
  res.status(200).send(family);
});

fastify.listen({ port: 3001 }, (err) => {
  if (err) throw err;
  fastify.log.info('Server started.');
});
