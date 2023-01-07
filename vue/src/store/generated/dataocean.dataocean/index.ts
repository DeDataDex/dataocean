import { Client, registry, MissingWalletError } from 'dataocean-client-ts'

import { Params } from "dataocean-client-ts/dataocean.dataocean/types"
import { Video } from "dataocean-client-ts/dataocean.dataocean/types"
import { VideoLink } from "dataocean-client-ts/dataocean.dataocean/types"


export { Params, Video, VideoLink };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				Video: {},
				VideoAll: {},
				VideoLink: {},
				VideoLinkAll: {},
				
				_Structure: {
						Params: getStructure(Params.fromPartial({})),
						Video: getStructure(Video.fromPartial({})),
						VideoLink: getStructure(VideoLink.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getVideo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Video[JSON.stringify(params)] ?? {}
		},
				getVideoAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VideoAll[JSON.stringify(params)] ?? {}
		},
				getVideoLink: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VideoLink[JSON.stringify(params)] ?? {}
		},
				getVideoLinkAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VideoLinkAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: dataocean.dataocean initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DataoceanDataocean.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVideo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DataoceanDataocean.query.queryVideo( key.id)).data
				
					
				commit('QUERY', { query: 'Video', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVideo', payload: { options: { all }, params: {...key},query }})
				return getters['getVideo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVideo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVideoAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DataoceanDataocean.query.queryVideoAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DataoceanDataocean.query.queryVideoAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'VideoAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVideoAll', payload: { options: { all }, params: {...key},query }})
				return getters['getVideoAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVideoAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVideoLink({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DataoceanDataocean.query.queryVideoLink( key.index)).data
				
					
				commit('QUERY', { query: 'VideoLink', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVideoLink', payload: { options: { all }, params: {...key},query }})
				return getters['getVideoLink']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVideoLink API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVideoLinkAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.DataoceanDataocean.query.queryVideoLinkAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.DataoceanDataocean.query.queryVideoLinkAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'VideoLinkAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVideoLinkAll', payload: { options: { all }, params: {...key},query }})
				return getters['getVideoLinkAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVideoLinkAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgPlayVideo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DataoceanDataocean.tx.sendMsgPlayVideo({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlayVideo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPlayVideo:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateVideo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DataoceanDataocean.tx.sendMsgCreateVideo({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVideo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVideo:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPaySign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DataoceanDataocean.tx.sendMsgPaySign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPaySign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPaySign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitPaySign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.DataoceanDataocean.tx.sendMsgSubmitPaySign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitPaySign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitPaySign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgPlayVideo({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DataoceanDataocean.tx.msgPlayVideo({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlayVideo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPlayVideo:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateVideo({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DataoceanDataocean.tx.msgCreateVideo({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVideo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVideo:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPaySign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DataoceanDataocean.tx.msgPaySign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPaySign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPaySign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitPaySign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.DataoceanDataocean.tx.msgSubmitPaySign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitPaySign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitPaySign:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
