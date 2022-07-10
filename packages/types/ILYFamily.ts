import { Static, Type } from '@sinclair/typebox';

export const Family = Type.Object({ yhw: Type.String(), lx: Type.String() });

type ILYFamily = Static<typeof Family>;

export default ILYFamily;
