import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPlayVideo } from "./types/dataocean/dataocean/tx";
import { MsgCreateVideo } from "./types/dataocean/dataocean/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dataocean.dataocean.MsgPlayVideo", MsgPlayVideo],
    ["/dataocean.dataocean.MsgCreateVideo", MsgCreateVideo],
    
];

export { msgTypes }