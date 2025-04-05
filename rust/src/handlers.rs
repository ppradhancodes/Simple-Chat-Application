use crate::models::{Message, User};
use crate::storage::Storage;
use uuid::Uuid;

pub struct ChatHandler {
    storage: Storage,
}

impl ChatHandler {
    pub fn new() -> Self {
        Self {
            storage: Storage::new(),
        }
    }

    pub fn register_or_login(&mut self, username: String) -> Result<User, String> {
        if let Some(existing_user) = self.storage.get_user_by_username(&username) {
            Ok(existing_user)
        } else {
            let user = User::new(username);
            self.storage.add_user(user.clone())?;
            Ok(user)
        }
    }

    pub fn send_message(
        &mut self,
        sender_id: Uuid,
        receiver_username: &str,
        content: String,
    ) -> Result<Message, String> {
        let receiver = self.storage
            .get_user_by_username(receiver_username)
            .ok_or_else(|| "Receiver not found".to_string())?;

        let message = Message::new(sender_id, receiver.id, content);
        self.storage.add_message(message.clone());
        Ok(message)
    }

    pub fn get_messages(&self, user_id: Uuid) -> Vec<Message> {
        self.storage.get_messages_for_user(&user_id)
    }

    pub fn search_messages(&self, keyword: &str, user_id: &Uuid) -> Vec<Message> {
        self.storage.search_messages(keyword, user_id)
    }

    pub fn delete_message(&mut self, keyword: &str, user_id: &Uuid) -> bool {
        self.storage.delete_message(keyword, user_id)
    }

    pub fn list_users(&self) -> Vec<User> {
        self.storage.list_users()
    }

    pub fn get_user(&self, id: &Uuid) -> Option<User> {
        self.storage.get_user(id)
    }
} 
