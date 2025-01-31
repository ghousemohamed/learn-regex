#!/bin/bash

# Add trap for Ctrl+C
trap 'echo -e "\nExiting tutorial..."; exit 0' INT

# Add progress tracking
PROGRESS_FILE="$HOME/.regex_tutorial_progress"

# Initialize lessons array
declare -A LESSONS=(
    ["1"]="Basic Patterns"
    ["2"]="Metacharacters"
    ["3"]="Character Classes"
    ["4"]="Quantifiers"
    ["5"]="Anchors"
    ["6"]="Word Boundaries"
    ["7"]="Groups and Capturing"
    ["8"]="Alternation"
    ["9"]="Greedy vs Lazy Quantifiers"
    ["10"]="Lookahead and Lookbehind"
    ["11"]="Character Classes Advanced"
    ["12"]="Backreferences"
    ["13"]="Final Practice"
)

# Function to display styled text
show_text() {
    echo "$1" | gum style --foreground 212
}

# Function to display headers
show_header() {
    gum style --border double --align center --width 50 --margin "1 2" --padding "1 2" "$1"
}

# Function to check regex match
check_match() {
    local pattern="$1"
    local test_string="$2"
    if [[ "$test_string" =~ $pattern ]]; then
        gum style --foreground 82 "✓ Correct! The pattern matches."
    else
        gum style --foreground 196 "✗ Try again. The pattern doesn't match."
    fi
}

# Function to show exit option
show_exit_option() {
    show_text "\nAt any time, press Ctrl+C to exit the tutorial"
    show_text "Or type 'exit' as your answer to quit\n"
}

# Function to show table of contents with progress
show_toc() {
    show_header "Table of Contents"
    
    # Create progress file if it doesn't exist
    touch "$PROGRESS_FILE"
    
    for num in "${!LESSONS[@]}"; do
        lesson="${LESSONS[$num]}"
        if grep -q "^$num$" "$PROGRESS_FILE"; then
            gum style --foreground 82 "[$num] ✓ $lesson"
        else
            gum style --foreground 244 "[$num] ○ $lesson"
        fi
    done
}

# Function to mark lesson as completed
complete_lesson() {
    local lesson_num="$1"
    if ! grep -q "^$lesson_num$" "$PROGRESS_FILE"; then
        echo "$lesson_num" >> "$PROGRESS_FILE"
    fi
}

# Function to get last completed lesson
get_last_lesson() {
    if [ -f "$PROGRESS_FILE" ]; then
        sort -n "$PROGRESS_FILE" | tail -n 1
    else
        echo "0"
    fi
}

# Function to start specific lesson
start_lesson() {
    local lesson_num="$1"
    case $lesson_num in
        1) lesson_basic_patterns ;;
        2) lesson_metacharacters ;;
        3) lesson_character_classes ;;
        4) lesson_quantifiers ;;
        5) lesson_anchors ;;
        6) lesson_word_boundaries ;;
        7) lesson_groups ;;
        8) lesson_alternation ;;
        9) lesson_greedy_lazy ;;
        10) lesson_lookaround ;;
        11) lesson_char_classes_advanced ;;
        12) lesson_backreferences ;;
        13) lesson_final_practice ;;
    esac
}

# Main tutorial
clear
show_header "Welcome to Interactive Regex Tutorial"

# Show table of contents and progress
show_toc

# Check for existing progress
LAST_LESSON=$(get_last_lesson)
if [ "$LAST_LESSON" -gt 0 ]; then
    NEXT_LESSON=$((LAST_LESSON + 1))
    if [ "$NEXT_LESSON" -le "${#LESSONS[@]}" ]; then
        show_text "You left off at lesson $LAST_LESSON: ${LESSONS[$LAST_LESSON]}"
        CHOICE=$(gum choose "Continue from lesson $NEXT_LESSON" "Start from beginning" "Exit")
    else
        CHOICE=$(gum choose "Start from beginning" "Exit")
    fi
else
    CHOICE=$(gum choose "Start from beginning" "Exit")
fi

case "$CHOICE" in
    "Continue from lesson"*)
        STARTING_LESSON=$NEXT_LESSON
        ;;
    "Start from beginning")
        STARTING_LESSON=1
        ;;
    "Exit")
        echo "Goodbye!"
        exit 0
        ;;
esac

# Split lessons into separate functions
lesson_basic_patterns() {
    clear
    show_header "Lesson 1: Basic Patterns"
    show_text "Let's start with simple character matching."
    show_text "Try to write a pattern that matches the word 'cat'"
    
    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "cat" ]]; then
            gum style --foreground 82 "Perfect! Simple patterns match exactly."
            complete_lesson "1"
            break
        else
            gum style --foreground 196 "Try again. We're looking for an exact match."
        fi
    done
}

lesson_metacharacters() {
    clear
    show_header "Lesson 2: Metacharacters"
    show_text "The dot (.) matches any single character."
    show_text "Write a pattern that matches 'cat', 'cot', and 'cut'"
    
    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "c.t" ]]; then
            gum style --foreground 82 "Excellent! The dot matches any character."
            complete_lesson "2"
            break
        else
            gum style --foreground 196 "Hint: Use . to match any single character"
        fi
    done
}

lesson_character_classes() {
    clear
    show_header "Lesson 3: Character Classes"
    show_text "Character classes [] match any single character within the brackets."
    show_text "Write a pattern that matches 'cat' and 'bat' but not 'rat'"
    
    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "[cb]at" ]]; then
            gum style --foreground 82 "Perfect! Character classes let you specify allowed characters."
            complete_lesson "3"
            break
        else
            gum style --foreground 196 "Hint: Use [cb] to match either 'c' or 'b'"
        fi
    done
}

lesson_quantifiers() {
    clear
    show_header "Lesson 4: Quantifiers"
    show_text "* means 'zero or more', + means 'one or more', ? means 'zero or one'"
    show_text "Write a pattern that matches 'color' and 'colour'"
    
    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "colou?r" ]]; then
            gum style --foreground 82 "Excellent! The ? makes the 'u' optional."
            complete_lesson "4"
            break
        else
            gum style --foreground 196 "Hint: Use ? to make a character optional"
        fi
    done
}

lesson_anchors() {
    clear
    show_header "Lesson 5: Anchors"
    show_text "Anchors (^ and $) match positions in text:"
    show_text "^ matches start of line"
    show_text "$ matches end of line"
    show_text "\nWrite a pattern that matches 'hello' only at the start of a line"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "^hello" ]]; then
            gum style --foreground 82 "Perfect! This will match 'hello' at line start."
            complete_lesson "5"
            break
        else
            gum style --foreground 196 "Hint: Use ^ to match the start of a line"
        fi
    done
}

lesson_word_boundaries() {
    clear
    show_header "Lesson 6: Word Boundaries"
    show_text "Word boundaries (\b) match positions between word and non-word characters"
    show_text "Write a pattern that matches 'cat' as a whole word only"
    show_text "(Should match 'cat' but not 'catch' or 'scatter')"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "\bcat\b" ]]; then
            gum style --foreground 82 "Excellent! This matches 'cat' as a whole word."
            complete_lesson "6"
            break
        else
            gum style --foreground 196 "Hint: Use \b at both ends of the word"
        fi
    done
}

lesson_groups() {
    clear
    show_header "Lesson 7: Groups and Capturing"
    show_text "Groups (parentheses) capture matched text and create sub-expressions"
    show_text "Write a pattern that matches and captures repeated words"
    show_text "Example: Should match 'the the' and capture 'the'"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "(\w+)\s+\1" ]]; then
            gum style --foreground 82 "Perfect! This captures and matches repeated words."
            complete_lesson "7"
            break
        else
            gum style --foreground 196 "Hint: Use () for capturing and \1 to reference the capture"
        fi
    done
}

lesson_alternation() {
    clear
    show_header "Lesson 8: Alternation"
    show_text "The | operator matches either of several patterns"
    show_text "Write a pattern that matches either 'cat' or 'dog' or 'bird'"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "cat|dog|bird" ]]; then
            gum style --foreground 82 "Excellent! This matches any of the three animals."
            complete_lesson "8"
            break
        else
            gum style --foreground 196 "Hint: Use | to separate alternatives"
        fi
    done
}

lesson_greedy_lazy() {
    clear
    show_header "Lesson 9: Greedy vs Lazy Quantifiers"
    show_text "By default, quantifiers are greedy (match as much as possible)"
    show_text "Adding ? makes them lazy (match as little as possible)"
    show_text "Write a pattern that lazily matches text between quotes"
    show_text "Example: \"hello\" \"world\" should match \"hello\" and \"world\" separately"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "\".*?\"" ]]; then
            gum style --foreground 82 "Perfect! This lazily matches quoted text."
            complete_lesson "9"
            break
        else
            gum style --foreground 196 "Hint: Use .*? for lazy matching"
        fi
    done
}

lesson_lookaround() {
    clear
    show_header "Lesson 10: Lookahead and Lookbehind"
    show_text "Lookaround assertions check for patterns without including them in the match"
    show_text "(?=...) positive lookahead"
    show_text "(?!...) negative lookahead"
    show_text "Write a pattern that matches numbers followed by 'px' but doesn't include 'px'"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "\d+(?=px)" ]]; then
            gum style --foreground 82 "Excellent! This matches numbers before 'px'."
            complete_lesson "10"
            break
        else
            gum style --foreground 196 "Hint: Use (?=...) to look ahead"
        fi
    done
}

lesson_char_classes_advanced() {
    clear
    show_header "Lesson 11: Character Classes Advanced"
    show_text "Character classes can include ranges and POSIX classes"
    show_text "Write a pattern that matches:"
    show_text "- Any uppercase letter A-Z"
    show_text "- Followed by exactly 3 digits"
    show_text "- Followed by any lowercase letter a-z"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "[A-Z]\d{3}[a-z]" ]]; then
            gum style --foreground 82 "Perfect! This uses character ranges effectively."
            complete_lesson "11"
            break
        else
            gum style --foreground 196 "Hint: Use [A-Z] for uppercase, \d{3} for digits, [a-z] for lowercase"
        fi
    done
}

lesson_backreferences() {
    clear
    show_header "Lesson 12: Backreferences"
    show_text "Backreferences (\1, \2, etc.) refer to previously captured groups"
    show_text "Write a pattern that matches HTML tags and their closing tags"
    show_text "Example: <b>text</b> or <i>text</i>"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "<([a-z]+)>.*?</\1>" ]]; then
            gum style --foreground 82 "Excellent! This matches paired HTML tags."
            complete_lesson "12"
            break
        else
            gum style --foreground 196 "Hint: Use () to capture the tag name and \1 to reference it"
        fi
    done
}

lesson_final_practice() {
    clear
    show_header "Lesson 13: Final Practice"
    show_text "Let's combine everything you've learned!"
    show_text "Write a pattern that matches:"
    show_text "- An email address"
    show_text "- Must start with letters/numbers/._-"
    show_text "- @ symbol"
    show_text "- Domain with letters/numbers"
    show_text "- .com, .org, or .net only"

    while true; do
        pattern=$(gum input --placeholder "Enter your regex pattern (or 'exit' to quit)")
        if [[ "$pattern" == "exit" ]]; then
            echo "Thanks for learning regex! Goodbye!"
            exit 0
        elif [[ "$pattern" == "^[a-zA-Z0-9._-]+@[a-zA-Z0-9]+\.(com|org|net)$" ]]; then
            gum style --foreground 82 "Congratulations! You've mastered regex!"
            complete_lesson "13"
            break
        else
            gum style --foreground 196 "Hint: Combine character classes, quantifiers, and alternation"
        fi
    done
}

# Start tutorial from chosen lesson
for ((i=STARTING_LESSON; i<=${#LESSONS[@]}; i++)); do
    start_lesson $i
    
    # Show progress after each lesson
    clear
    show_toc
    
    if [ $i -lt ${#LESSONS[@]} ]; then
        gum confirm "Continue to next lesson?" || exit 0
    fi
done

# Final message
clear
show_header "🎉 Congratulations!"
show_text "You've completed the basic regex tutorial!"
show_text "Practice makes perfect - keep experimenting with patterns!"

gum confirm "Would you like to try more exercises?" && exec $0 || exit 0 