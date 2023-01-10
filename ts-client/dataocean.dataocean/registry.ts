import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSubmitPaySign } from "./types/dataocean/dataocean/tx";
import { MsgCreateVideo } from "./types/dataocean/dataocean/tx";
import { MsgPlayVideo } from "./types/dataocean/dataocean/tx";
import { MsgPaySign } from "./types/dataocean/dataocean/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dataocean.dataocean.MsgSubmitPaySign", MsgSubmitPaySign],
    ["/dataocean.dataocean.MsgCreateVideo", MsgCreateVideo],
    ["/dataocean.dataocean.MsgPlayVideo", MsgPlayVideo],
    ["/dataocean.dataocean.MsgPaySign", MsgPaySign],
    
];

export { msgTypes }