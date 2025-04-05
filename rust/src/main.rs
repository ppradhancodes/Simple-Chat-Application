mod models;
mod storage;
mod handlers;
mod utils;

use handlers::ChatHandler;
use std::io::{self, Write};
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    println!("Simple Chat Application");
    let mut chat_handler = ChatHandler::new();
    let mut current_user = None;

    loop {
        if current_user.is_none() {
            print!("Enter username (or 'exit' to quit): ");
            io::stdout().flush()?;
            let mut username = String::new();
            io::stdin().read_line(&mut username)?;
            let username = username.trim();

            if username == "exit" {
                break;
            }

            match chat_handler.register_or_login(username.to_string()) {
                Ok(user) => {
                    println!("Welcome, {}!", username);
                    current_user = Some(user);
                }
                Err(e) => println!("Error: {}", e),
            }
        } else {
            print!("\nCommands:\n1. Send message\n2. View messages\n3. Search messages\n4. List users\n5. Logout\nChoice: ");
            io::stdout().flush()?;
            let mut choice = String::new();
            io::stdin().read_line(&mut choice)?;

            match choice.trim() {
                "1" => {
                    print!("Enter recipient username: ");
                    io::stdout().flush()?;
                    let mut recipient = String::new();
                    io::stdin().read_line(&mut recipient)?;
                    
                    print!("Enter message: ");
                    io::stdout().flush()?;
                    let mut content = String::new();
                    io::stdin().read_line(&mut content)?;

                    match chat_handler.send_message(
                        current_user.as_ref().unwrap().id,
                        recipient.trim(),
                        content.trim().to_string(),
                    ) {
                        Ok(_) => println!("Message sent!"),
                        Err(e) => println!("Error: {}", e),
                    }
                }
                "2" => {
                    let messages = chat_handler.get_messages(current_user.as_ref().unwrap().id);
                    if messages.is_empty() {
                        println!("No messages.");
                    } else {
                        for msg in messages {
                            let sender = chat_handler.get_user(&msg.sender_id);
                            let sender_name = sender.map(|u| u.username).unwrap_or_else(|| "Unknown".to_string());
                            utils::print_message(&sender_name, &msg.content, msg.timestamp);
                        }
                    }
                }
                "3" => {
                    print!("Enter search keyword: ");
                    io::stdout().flush()?;
                    let mut keyword = String::new();
                    io::stdin().read_line(&mut keyword)?;

                    let messages = chat_handler.search_messages(
                        keyword.trim(),
                        &current_user.as_ref().unwrap().id
                    );
                    if messages.is_empty() {
                        println!("No messages found.");
                    } else {
                        for msg in messages {
                            let sender = chat_handler.get_user(&msg.sender_id);
                            let sender_name = sender.map(|u| u.username).unwrap_or_else(|| "Unknown".to_string());
                            utils::print_message(&sender_name, &msg.content, msg.timestamp);
                        }
                    }
                }
                "4" => {
                    let users = chat_handler.list_users();
                    println!("\nRegistered users:");
                    for user in users {
                        println!("- {}", user.username);
                    }
                }
                "5" => {
                    println!("Goodbye, {}!", current_user.as_ref().unwrap().username);
                    current_user = None;
                }
                _ => println!("Invalid choice"),
            }
        }
    }

    Ok(())
}
