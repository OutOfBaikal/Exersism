use std::collections::HashMap;

pub fn plants(diagram: &str, student: &str) -> Vec<&'static str> {
    let plant_map: HashMap<char, &'static str> = [
        ('G', "grass"),
        ('C', "clover"),
        ('R', "radishes"),
        ('V', "violets"),
    ].iter().cloned().collect();
    
    let students = [
        "Alice", "Bob", "Charlie", "David", "Eve", "Fred",
        "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry",
    ];

    let student_index = match students.iter().position(|s| *s == student) {
        Some(idx) => idx,
        None => panic!("Unknown student: {}", student),
    };
    
    let lines: Vec<&str> = diagram
        .lines()
        .filter(|line| !line.trim().is_empty())
        .collect();
    if lines.len() < 2 {
        panic!("Diagram must have at least 2 rows");
    }

    let top_row = lines[0].trim();
    let bottom_row = lines[1].trim();
    
    let start_pos = student_index * 2;
    if start_pos + 1 >= top_row.len() || start_pos + 1 >= bottom_row.len() {
        panic!("Not enough plants for student: {}", student);
    }

    let chars = [
        top_row.chars().nth(start_pos).unwrap(),
        top_row.chars().nth(start_pos + 1).unwrap(),
        bottom_row.chars().nth(start_pos).unwrap(),
        bottom_row.chars().nth(start_pos + 1).unwrap(),
    ];

    chars.iter()
        .map(|&c| *plant_map.get(&c).expect("Invalid plant code"))
        .collect()
}
