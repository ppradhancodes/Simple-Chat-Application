use crate::models::{Message, User};
use std::collections::HashMap;
use uuid::Uuid;

#[derive(Default)]
pub struct Storage {
    users: HashMap<Uuid, User>,
    messages: Vec<Message>,
}

impl Storage {
    pub fn new() -> Self {
        Self {
            users: HashMap::new(),
            messages: Vec::new(),
        }
    }

    pub fn add_user(&mut self, user: User) -> Result<(), String> {
        if self.users.values().any(|u| u.username == user.username) {
            return Err("Username already exists".to_string());
        }
        self.users.insert(user.id, user);
        Ok(())
    }

    pub fn get_user(&self, id: &Uuid) -> Option<User> {
        self.users.get(id).cloned()
    }

    pub fn get_user_by_username(&self, username: &str) -> Option<User> {
        self.users
            .values()
            .find(|u| u.username == username)
            .cloned()
    }

    pub fn add_message(&mut self, message: Message) {
        self.messages.push(message);
    }

    pub fn get_messages_for_user(&self, user_id: &Uuid) -> Vec<Message> {
        self.messages
            .iter()
            .filter(|m| m.sender_id == *user_id || m.receiver_id == *user_id)
            .cloned()
            .collect()
    }

    pub fn search_messages(&self, keyword: &str, user_id: &Uuid) -> Vec<Message> {
        self.messages
            .iter()
            .filter(|m| m.sender_id == *user_id || m.receiver_id == *user_id)
            .filter(|m| m.content.to_lowercase().contains(&keyword.to_lowercase()))
            .cloned()
            .collect()
    }

    pub fn delete_message(&mut self, keyword: &str, user_id: &Uuid) -> bool {
        let non_delete_messages: Vec<Message>  = self.messages.iter().filter(|m| (m.sender_id != *user_id && m.receiver_id != *user_id) || !m.content.to_lowercase().contains(&keyword.to_lowercase())).map(|m| m.clone()).collect();
        let is_delete: bool = non_delete_messages.len() != self.messages.len();
        self.messages = non_delete_messages;
        return is_delete;
    }

    pub fn list_users(&self) -> Vec<User> {
        self.users.values().cloned().collect()
    }
} 
