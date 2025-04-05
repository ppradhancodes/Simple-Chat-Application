use chrono::{DateTime, Utc};

pub fn format_timestamp(timestamp: DateTime<Utc>) -> String {
    timestamp.format("%Y-%m-%d %H:%M:%S").to_string()
}

pub fn print_message(username: &str, content: &str, timestamp: DateTime<Utc>) {
    println!(
        "[{}] {}: {}",
        format_timestamp(timestamp),
        username,
        content
    );
} 