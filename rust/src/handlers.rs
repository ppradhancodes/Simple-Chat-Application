use crate::models::{Message, User};
use crate::storage::Storage;
use std::sync::{Arc, Mutex};
use uuid::Uuid;

pub struct ChatHandler {
    storage: Arc<Mutex<Storage>>,
}

impl ChatHandler {
    pub fn new() -> Self {
        Self {
            storage: Storage::new(),
        }
    }

    pub fn register_or_login(&self, username: String) -> Result<User, String> {
        let storage = self.storage.lock().unwrap();
        // Check if user already exists
        if let Some(existing_user) = storage.get_user_by_username(&username) {
            // User exists, return the existing user (login)
            Ok(existing_user)
        } else {
            // User doesn't exist, create new user (register)
            drop(storage); // Release the read lock before getting write lock
            let user = User::new(username);
            let mut storage = self.storage.lock().unwrap();
            storage.add_user(user.clone())?;
            Ok(user)
        }
    }

    pub fn send_message(
        &self,
        sender_id: Uuid,
        receiver_username: &str,
        content: String,
    ) -> Result<Message, String> {
        let storage = self.storage.lock().unwrap();
        let receiver = storage
            .get_user_by_username(receiver_username)
            .ok_or_else(|| "Receiver not found".to_string())?;

        let message = Message::new(sender_id, receiver.id, content);
        drop(storage);

        let mut storage = self.storage.lock().unwrap();
        storage.add_message(message.clone());
        Ok(message)
    }

    pub fn get_messages(&self, user_id: Uuid) -> Vec<Message> {
        let storage = self.storage.lock().unwrap();
        storage.get_messages_for_user(&user_id)
    }

    pub fn search_messages(&self, keyword: &str) -> Vec<Message> {
        let storage = self.storage.lock().unwrap();
        storage.search_messages(keyword)
    }

    pub fn list_users(&self) -> Vec<User> {
        let storage = self.storage.lock().unwrap();
        storage.list_users()
    }

    pub fn get_user(&self, id: &Uuid) -> Option<User> {
        let storage = self.storage.lock().unwrap();
        storage.get_user(id)
    }
} 