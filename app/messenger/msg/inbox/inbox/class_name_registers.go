/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2023-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package inbox

const (
	Predicate_inboxMessageData                  = "inboxMessageData"
	Predicate_inboxMessageId                    = "inboxMessageId"
	Predicate_inbox_sendUserMessageToInbox      = "inbox_sendUserMessageToInbox"
	Predicate_inbox_sendChatMessageToInbox      = "inbox_sendChatMessageToInbox"
	Predicate_inbox_sendUserMultiMessageToInbox = "inbox_sendUserMultiMessageToInbox"
	Predicate_inbox_sendChatMultiMessageToInbox = "inbox_sendChatMultiMessageToInbox"
	Predicate_inbox_editUserMessageToInbox      = "inbox_editUserMessageToInbox"
	Predicate_inbox_editChatMessageToInbox      = "inbox_editChatMessageToInbox"
	Predicate_inbox_deleteMessagesToInbox       = "inbox_deleteMessagesToInbox"
	Predicate_inbox_deleteUserHistoryToInbox    = "inbox_deleteUserHistoryToInbox"
	Predicate_inbox_deleteChatHistoryToInbox    = "inbox_deleteChatHistoryToInbox"
	Predicate_inbox_readUserMediaUnreadToInbox  = "inbox_readUserMediaUnreadToInbox"
	Predicate_inbox_readChatMediaUnreadToInbox  = "inbox_readChatMediaUnreadToInbox"
	Predicate_inbox_updateHistoryReaded         = "inbox_updateHistoryReaded"
	Predicate_inbox_updatePinnedMessage         = "inbox_updatePinnedMessage"
	Predicate_inbox_unpinAllMessages            = "inbox_unpinAllMessages"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_inboxMessageData: {
		0: 1002286548, // 0x3bbdadd4

	},
	Predicate_inboxMessageId: {
		0: -963460705, // 0xc692c19f

	},
	Predicate_inbox_sendUserMessageToInbox: {
		0: -208741709, // 0xf38edab3

	},
	Predicate_inbox_sendChatMessageToInbox: {
		0: -1760197438, // 0x971584c2

	},
	Predicate_inbox_sendUserMultiMessageToInbox: {
		0: -1782288007, // 0x95c47179

	},
	Predicate_inbox_sendChatMultiMessageToInbox: {
		0: -694455924, // 0xd69b718c

	},
	Predicate_inbox_editUserMessageToInbox: {
		0: 1559967656, // 0x5cfb37a8

	},
	Predicate_inbox_editChatMessageToInbox: {
		0: 2031122959, // 0x79107a0f

	},
	Predicate_inbox_deleteMessagesToInbox: {
		0: -2061734348, // 0x851c6e34

	},
	Predicate_inbox_deleteUserHistoryToInbox: {
		0: 336232792, // 0x140a8158

	},
	Predicate_inbox_deleteChatHistoryToInbox: {
		0: -659905022, // 0xd8aaa602

	},
	Predicate_inbox_readUserMediaUnreadToInbox: {
		0: 364970827, // 0x15c1034b

	},
	Predicate_inbox_readChatMediaUnreadToInbox: {
		0: 1430347220, // 0x55415dd4

	},
	Predicate_inbox_updateHistoryReaded: {
		0: -1010283296, // 0xc3c84ce0

	},
	Predicate_inbox_updatePinnedMessage: {
		0: -1452528908, // 0xa96c2af4

	},
	Predicate_inbox_unpinAllMessages: {
		0: 589079137, // 0x231ca261

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	1002286548:  Predicate_inboxMessageData,                  // 0x3bbdadd4
	-963460705:  Predicate_inboxMessageId,                    // 0xc692c19f
	-208741709:  Predicate_inbox_sendUserMessageToInbox,      // 0xf38edab3
	-1760197438: Predicate_inbox_sendChatMessageToInbox,      // 0x971584c2
	-1782288007: Predicate_inbox_sendUserMultiMessageToInbox, // 0x95c47179
	-694455924:  Predicate_inbox_sendChatMultiMessageToInbox, // 0xd69b718c
	1559967656:  Predicate_inbox_editUserMessageToInbox,      // 0x5cfb37a8
	2031122959:  Predicate_inbox_editChatMessageToInbox,      // 0x79107a0f
	-2061734348: Predicate_inbox_deleteMessagesToInbox,       // 0x851c6e34
	336232792:   Predicate_inbox_deleteUserHistoryToInbox,    // 0x140a8158
	-659905022:  Predicate_inbox_deleteChatHistoryToInbox,    // 0xd8aaa602
	364970827:   Predicate_inbox_readUserMediaUnreadToInbox,  // 0x15c1034b
	1430347220:  Predicate_inbox_readChatMediaUnreadToInbox,  // 0x55415dd4
	-1010283296: Predicate_inbox_updateHistoryReaded,         // 0xc3c84ce0
	-1452528908: Predicate_inbox_updatePinnedMessage,         // 0xa96c2af4
	589079137:   Predicate_inbox_unpinAllMessages,            // 0x231ca261

}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}
