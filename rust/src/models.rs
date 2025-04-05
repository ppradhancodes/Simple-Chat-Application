use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct User {
    pub id: Uuid,
    pub username: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Message {
    pub id: Uuid,
    pub sender_id: Uuid,
    pub receiver_id: Uuid,
    pub content: String,
    pub timestamp: DateTime<Utc>,
}

impl User {
    pub fn new(username: String) -> Self {
        Self {
            id: Uuid::new_v4(),
            username,
        }
    }
}

impl Message {
    pub fn new(sender_id: Uuid, receiver_id: Uuid, content: String) -> Self {
        Self {
            id: Uuid::new_v4(),
            sender_id,
            receiver_id,
            content,
            timestamp: Utc::now(),
        }
    }
} 