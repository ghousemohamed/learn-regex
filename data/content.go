package data

import "github.com/ghousemohamed/regex-in-the-terminal/models"

func GetLessons() []models.Lesson {
	return []models.Lesson{
		{
			Title:       "Basic Patterns",
			Description: "Welcome to your first regex lesson! Let's start with the simplest concept: literal matching.\n\n" +
				"In regex, when you type normal letters or numbers, they match exactly what you type. It's like a simple search function.\n\n" +
				"For example:\n" +
				"- 'dog' matches exactly 'dog'\n" +
				"- '123' matches exactly '123'\n" +
				"- 'hello' matches exactly 'hello'\n\n" +
				"However, be careful! Exact matching is strict:\n" +
				"- It won't match if there are extra characters\n" +
				"- It won't match if the case is different\n" +
				"- It won't match if it's part of another word\n\n" +
				"Test text:\n" +
				"✓ cat (exact match)\n" +
				"✗ bats \n" +
				"✗ Cat (different case)\n" +
				"✗ catch (part of another word)",
			Task:        "Write a pattern that matches exactly the word 'cat'",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "bats", Expected: false},
				{Text: "Cat food", Expected: false},
				{Text: "dog", Expected: false},
			},
		},
		{
			Title:       "The Dot Metacharacter",
			Description: "Now let's learn about our first special character: the dot (.)\n\n" +
				"The dot is a wildcard character that matches any single character except a newline. Think of it as a placeholder " +
				"that says 'I don't care what character is here, as long as there is one'.\n\n" +
				"Examples:\n" +
				"- 'h.t' matches 'hat', 'hot', 'hit'\n" +
				"- 'b.g' matches 'bag', 'big', 'bug'\n" +
				"- 'r.n' matches 'run', 'ran', 'ryn'\n\n" +
				"Remember: The dot matches EXACTLY ONE character - no more, no less!\n\n" +
				"Test text:\n" +
				"✓ cat (single character between 'c' and 't')\n" +
				"✓ cot (single character between 'c' and 't')\n" +
				"✓ cut (single character between 'c' and 't')\n" +
				"✗ cart (two characters between 'c' and 't')\n" +
				"✗ ct (no character between 'c' and 't')",
			Task:        "Write a pattern that matches 'cat', 'cot', and 'cut'",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "cot", Expected: true},
				{Text: "cut", Expected: true},
				{Text: "cart", Expected: false},
				{Text: "ct", Expected: false},
			},
		},
		{
			Title:       "Simple Character Classes",
			Description: "Let's learn about character classes - one of the most powerful features in regex!\n\n" +
				"A character class is created with square brackets [] and matches ANY SINGLE character from the set you specify inside the brackets.\n\n" +
				"Examples:\n" +
				"- [aeiou] matches any single vowel\n" +
				"- [RGB] matches either 'R', 'G', or 'B'\n" +
				"- [123] matches '1', '2', or '3'\n\n" +
				"This is super useful when you want to:\n" +
				"- Match multiple possible characters at a specific position\n" +
				"- Allow for variations in spelling\n" +
				"- Match patterns with alternative characters\n\n" +
				"Test text:\n" +
				"✓ cat (starts with 'c')\n" +
				"✓ bat (starts with 'b')\n" +
				"✗ rat (starts with 'r')\n" +
				"✗ mat (starts with 'm')",
			Task:        "Write a pattern that matches both 'cat' and 'bat' but not 'rat'",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "bat", Expected: true},
				{Text: "rat", Expected: false},
			},
		},
		{
			Title:       "Negated Character Classes",
			Description: "Let's flip character classes on their head with negation!\n\n" +
				"When you put a ^ as the first character inside brackets [^...], it matches any character that is NOT in the set. " +
				"Think of it as saying 'match anything EXCEPT these characters'.\n\n" +
				"Examples:\n" +
				"- [^0-9] matches any non-digit character\n" +
				"- [^aeiou] matches any non-vowel\n" +
				"- [^xyz] matches anything except x, y, or z\n\n" +
				"Remember: The ^ only means negation when it's the first character inside the brackets. " +
				"Anywhere else in the brackets, ^ just matches a literal ^ character.\n\n" +
				"Test text:\n" +
				"✓ cat (starts with 'c')\n" +
				"✓ bat (starts with 'b')\n" +
				"✗ rat (starts with 'r')\n" +
				"✗ mat (starts with 'm')",
			Task:        "Write a pattern that matches 'cat' and 'bat' but NOT 'rat' or 'mat' using negation",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "bat", Expected: true},
				{Text: "rat", Expected: false},
				{Text: "mat", Expected: false},
			},
		},
		{
			Title:       "Character Ranges",
			Description: "Let's learn about a shortcut for character classes - ranges!\n\n" +
				"Instead of listing every character you want to match, you can use a hyphen (-) between two characters to match any single character in that range.\n\n" +
				"Common ranges:\n" +
				"- [1-5] matches any single digit from 1 to 5\n" +
				"- [m-p] matches any single letter from m to p\n" +
				"- [D-G] matches any single uppercase letter from D to G\n\n" +
				"This is particularly useful for:\n" +
				"- Matching ranges of numbers\n" +
				"- Matching ranges of letters\n" +
				"- Creating more concise patterns\n\n" +
				"Pro tip: The order matters in ranges - the first character must come before the second in ASCII order!\n\n" +
				"Test text:\n" +
				"✓ dog (three lowercase letters)\n" +
				"✓ cat (three lowercase letters)\n" +
				"✗ DOG (uppercase letters)\n" +
				"✗ d0g (contains a number)",
			Task:        "Write a pattern that matches any three-letter word using lowercase letters",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "dog", Expected: true},
				{Text: "bat", Expected: true},
				{Text: "rat", Expected: true},
				{Text: "ct", Expected: false},
			},
		},
		{
			Title:       "Multiple Ranges",
			Description: "Let's combine ranges to create more powerful patterns!\n\n" +
				"You can put multiple ranges inside the same character class to match characters from any of those ranges.\n\n" +
				"Examples:\n" +
				"- [1-3A-C] matches 1, 2, 3, A, B, or C\n" +
				"- [b-dx-z] matches b, c, d, x, y, or z\n" +
				"- [0-4a-c] matches 0, 1, 2, 3, 4, a, b, or c\n\n" +
				"This is particularly useful for:\n" +
				"- Matching mixed character types\n" +
				"- Creating flexible patterns\n" +
				"- Handling multiple valid ranges\n\n" +
				"Test text:\n" +
				"✓ P8 (uppercase letter + digit)\n" +
				"✓ m4 (lowercase letter + digit)\n" +
				"✗ 8P (digit + letter)\n" +
				"✗ mp (no digit)\n" +
				"✗ p88 (too many digits)",
			Task:        "Write a pattern that matches words starting with any letter (upper or lower) followed by two digits",
			TestCases: []models.TestCase{
				{Text: "A12", Expected: true},
				{Text: "b45", Expected: true},
				{Text: "Z90", Expected: true},
				{Text: "123", Expected: false},
				{Text: "abc", Expected: false},
				{Text: "A1", Expected: false},
				{Text: "ABC", Expected: false},
			},
		},
		{
			Title:       "Optional Characters",
			Description: "Let's learn about making characters optional!\n\n" +
				"The question mark (?) makes the character before it optional - meaning it can appear once or not at all.\n\n" +
				"Examples:\n" +
				"- 'files?' matches 'file' and 'files'\n" +
				"- 'Nov(ember)?' matches 'Nov' and 'November'\n" +
				"- 'https?' matches 'http' and 'https'\n\n" +
				"This is particularly useful for:\n" +
				"- Handling optional suffixes\n" +
				"- Matching different spellings\n" +
				"- Making parts of a pattern optional\n\n" +
				"Test text:\n" +
				"✓ color (American spelling)\n" +
				"✓ colour (British spelling)\n" +
				"✗ colouur (too many u's)\n" +
				"✗ culor (wrong vowel)",
			Task:        "Write a pattern that matches both 'color' and 'colour'",
			TestCases: []models.TestCase{
				{Text: "color", Expected: true},
				{Text: "colour", Expected: true},
				{Text: "colouur", Expected: false},
				{Text: "colouurr", Expected: false},
			},
		},
		{
			Title:       "Zero or More",
			Description: "Let's learn about the asterisk (*) - a powerful repetition operator!\n\n" +
				"The asterisk (*) means 'zero or more occurrences' of the character before it.\n\n" +
				"Examples:\n" +
				"- 'bo*m' matches 'bm', 'bom', 'boom', 'booom'\n" +
				"- 'he*y' matches 'hy', 'hey', 'heey', 'heeey'\n" +
				"- 'w*in' matches 'in', 'win', 'wwwin'\n\n" +
				"This is particularly useful for:\n" +
				"- Matching repeated characters\n" +
				"- Making parts of a pattern completely optional\n" +
				"- Handling variable-length patterns\n\n" +
				"Test text:\n" +
				"✓ ca (no t's)\n" +
				"✓ cat (one t)\n" +
				"✓ catt (two t's)\n" +
				"✓ cattt (three t's)\n" +
				"✗ ct (missing 'a')",
			Task:        "Write a pattern that matches 'ca' followed by any number of 't's (including none)",
			TestCases: []models.TestCase{
				{Text: "ca", Expected: true},
				{Text: "cat", Expected: true},
				{Text: "catt", Expected: true},
				{Text: "cattt", Expected: true},
				{Text: "ct", Expected: false},
			},
		},
		{
			Title:       "One or More",
			Description: "Let's learn about the plus sign (+) - a quantifier that ensures something appears!\n\n" +
				"The plus sign (+) means 'one or more occurrences' of the character before it. Unlike *, it requires at least one match.\n\n" +
				"Examples:\n" +
				"- 'ho+p' matches 'hop', 'hoop', 'hooop', but not 'hp'\n" +
				"- 'wa+ve' matches 'wave', 'waave', 'waaave', but not 'wve'\n" +
				"- '[0-9]+' matches numbers of any length (but not empty)\n\n" +
				"This is particularly useful for:\n" +
				"- Ensuring at least one occurrence\n" +
				"- Matching non-empty sequences\n" +
				"- Required repetitions\n\n" +
				"Test text:\n" +
				"✓ cat (one t)\n" +
				"✓ catt (two t's)\n" +
				"✓ cattt (three t's)\n" +
				"✗ ca (no t's)",
			Task:        "Write a pattern that matches 'cat' with one or more t's",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "catt", Expected: true},
				{Text: "cattt", Expected: true},
				{Text: "ca", Expected: false},
			},
		},
		{
			Title:       "Exact Count",
			Description: "Let's learn about precise repetition with curly braces!\n\n" +
				"The {n} syntax specifies exactly n occurrences of the previous character.\n\n" +
				"Examples:\n" +
				"- 'w{3}' matches exactly three w's ('www')\n" +
				"- '[0-9]{2}' matches exactly two digits\n" +
				"- 'hi{3}' matches 'hiii' only\n\n" +
				"This is particularly useful for:\n" +
				"- Fixed-length codes\n" +
				"- Exact pattern lengths\n" +
				"- Precise matching\n\n" +
				"Test text:\n" +
				"✓ aaab (three a's then b)\n" +
				"✗ aab (too few a's)\n" +
				"✗ aaaab (too many a's)\n" +
				"✗ ab (too few a's)",
			Task:        "Write a pattern that matches exactly three 'a's followed by 'b' (and nothing else)",
			TestCases: []models.TestCase{
				{Text: "aaab", Expected: true},
				{Text: "aab", Expected: false},
				{Text: "aaaab", Expected: false},
				{Text: "ab", Expected: false},
			},
		},
		{
			Title:       "Range of Counts",
			Description: "Let's learn about flexible repetition ranges!\n\n" +
				"The {min,max} syntax allows a character to repeat between min and max times.\n\n" +
				"Examples:\n" +
				"- 'x{1,3}' matches 'x', 'xx', or 'xxx'\n" +
				"- '[0-9]{2,4}' matches numbers with 2-4 digits\n" +
				"- 'hi{0,2}' matches 'h', 'hi', or 'hii'\n\n" +
				"You can also use:\n" +
				"- {2,} for 2 or more\n" +
				"- {,3} for up to 3\n\n" +
				"Test text:\n" +
				"✗ ab (too few b's)\n" +
				"✓ abb (two b's)\n" +
				"✓ abbb (three b's)\n" +
				"✓ abbbb (four b's)\n" +
				"✗ abbbbb (too many b's)",
			Task:        "Write a pattern that matches 'ab' followed by 2 to 4 'b's",
			TestCases: []models.TestCase{
				{Text: "ab", Expected: false},
				{Text: "abb", Expected: true},
				{Text: "abbb", Expected: true},
				{Text: "abbbb", Expected: true},
			},
		},
		{
			Title:       "Start Anchor",
			Description: "Let's learn about position matching with the caret (^)!\n\n" +
				"The caret (^) matches the start of a line when used outside of square brackets.\n\n" +
				"Examples:\n" +
				"- '^start' matches 'start here' but not 'the start'\n" +
				"- '^[0-9]' matches lines beginning with a digit\n" +
				"- '^>' matches lines starting with '>'\n\n" +
				"This is particularly useful for:\n" +
				"- Line beginnings\n" +
				"- String validation\n" +
				"- Pattern positioning\n\n" +
				"Test text:\n" +
				"✓ hello (at start of line)\n" +
				"✓ hello world (at start of line)\n" +
				"✗ world hello (in middle)\n" +
				"✗ say hello (at end)",
			Task:        "Write a pattern that matches 'hello' only at the start of a line",
			TestCases: []models.TestCase{
				{Text: "hello", Expected: true},
				{Text: "hello world", Expected: true},
				{Text: "world hello", Expected: false},
				{Text: "say hello", Expected: false},
			},
		},
		{
			Title:       "End Anchor",
			Description: "Let's learn about matching at the end of lines!\n\n" +
				"The dollar sign ($) matches the position at the end of a line. It ensures a pattern appears at the end.\n\n" +
				"Examples:\n" +
				"- 'bye$' matches 'goodbye' but not 'bye now'\n" +
				"- '[0-9]$' matches '123' but not '123a'\n" +
				"- 'day$' matches 'today' but not 'days'\n\n" +
				"This is particularly useful for:\n" +
				"- Validating string endings\n" +
				"- Matching final words\n" +
				"- Ensuring nothing follows\n\n" +
				"Test text:\n" +
				"✓ world (ends with 'world')\n" +
				"✓ hello world (ends with 'world')\n" +
				"✗ world hello (doesn't end with 'world')\n" +
				"✗ world now (doesn't end with 'world')",
			Task:        "Write a pattern that matches 'world' only at the end of a line",
			TestCases: []models.TestCase{
				{Text: "world", Expected: true},
				{Text: "hello world", Expected: true},
				{Text: "world hello", Expected: false},
				{Text: "world now", Expected: false},
			},
		},
		{
			Title:       "Word Boundaries",
			Description: "Let's learn about finding complete words!\n\n" +
				"Word boundaries (\\b) match positions where a word character (letter, number, underscore) meets a non-word character.\n\n" +
				"Examples:\n" +
				"- '\\brun\\b' matches 'go run now' but not 'running'\n" +
				"- '\\bthe\\b' matches 'see the cat' but not 'there'\n" +
				"- '\\bup\\b' matches 'stand up tall' but not 'upper'\n\n" +
				"This is particularly useful for:\n" +
				"- Finding whole words\n" +
				"- Avoiding partial matches\n" +
				"- Precise word matching\n\n" +
				"Test text:\n" +
				"✓ cat (complete word)\n" +
				"✗ cats (part of longer word)\n" +
				"✗ scatter (contains 'cat' inside)\n" +
				"✓ cat food (complete word)",
			Task:        "Write a pattern that matches 'cat' as a complete word only",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "cats", Expected: false},
				{Text: "scatter", Expected: false},
				{Text: "cat food", Expected: true},
			},
		},
		{
			Title:       "Grouping",
			Description: "Let's learn about grouping patterns together!\n\n" +
				"Parentheses () let you treat multiple characters as a single unit and apply operations to them together.\n\n" +
				"Examples:\n" +
				"- '(hi){2}' matches 'hihi'\n" +
				"- '(ab)+' matches 'ab', 'abab', 'ababab'\n" +
				"- '(good|nice)' matches 'good' or 'nice'\n\n" +
				"This is particularly useful for:\n" +
				"- Repeating sequences\n" +
				"- Applying quantifiers\n" +
				"- Creating sub-patterns\n\n" +
				"Test text:\n" +
				"✗ ha (single occurrence)\n" +
				"✓ haha (exactly two occurrences)\n" +
				"✗ hahaha (too many occurrences)\n" +
				"✗ ah (wrong order)",
			Task:        "Write a pattern that matches 'ha' repeated exactly twice",
			TestCases: []models.TestCase{
				{Text: "ha", Expected: false},
				{Text: "haha", Expected: true},
				{Text: "hahaha", Expected: false},
				{Text: "ah", Expected: false},
			},
		},
		{
			Title:       "Alternation",
			Description: "Let's learn about matching alternatives!\n\n" +
				"The vertical bar (|) lets you match one pattern OR another pattern.\n\n" +
				"Examples:\n" +
				"- 'yes|no' matches either 'yes' or 'no'\n" +
				"- 'hi|hello|hey' matches any of these greetings\n" +
				"- 'Mon|Tue|Wed' matches any of these days\n\n" +
				"This is particularly useful for:\n" +
				"- Multiple options\n" +
				"- Alternative patterns\n" +
				"- Choice matching\n\n" +
				"Test text:\n" +
				"✓ cat (first option)\n" +
				"✓ dog (second option)\n" +
				"✗ catdog (not a single word)\n" +
				"✗ bird (not in options)",
			Task:        "Write a pattern that matches either 'cat' or 'dog' as complete words",
			TestCases: []models.TestCase{
				{Text: "cat", Expected: true},
				{Text: "dog", Expected: true},
				{Text: "catdog", Expected: false},
				{Text: "mouse", Expected: false},
			},
		},
		{
			Title:       "Common Shortcuts",
			Description: "Let's learn about regex shorthand characters!\n\n" +
				"Instead of writing long character classes, regex provides convenient shortcuts:\n" +
				"- \\d matches any digit\n" +
				"- \\w matches any word character (letters, numbers, underscore)\n" +
				"- \\s matches any whitespace\n\n" +
				"Examples:\n" +
				"- 'age: \\d' matches 'age: 5', 'age: 7', etc.\n" +
				"- '\\w_\\w' matches 'a_b', 'x_y', '1_2', etc.\n" +
				"- 'hi\\sworld' matches 'hi world'\n\n" +
				"This is particularly useful for:\n" +
				"- Matching common character types\n" +
				"- Writing cleaner patterns\n" +
				"- Quick character class shortcuts\n\n" +
				"Test text:\n" +
				"✓ a1 (letter followed by number)\n" +
				"✓ x9 (letter followed by number)\n" +
				"✗ 1a (number followed by letter)\n" +
				"✗ ab (no number)\n" +
				"✗ 12 (no letter)",
			Task:        "Write a pattern that matches a word character followed by a digit",
			TestCases: []models.TestCase{
				{Text: "a1", Expected: true},
				{Text: "x9", Expected: true},
				{Text: "1a", Expected: false},
				{Text: "ab", Expected: false},
				{Text: "12", Expected: false},
			},
		},
		{
			Title:       "Non-Capturing Groups",
			Description: "Let's learn about a special kind of grouping!\n\n" +
				"Sometimes we want to group patterns but don't need to remember what they matched. " +
				"Non-capturing groups (?:pattern) do exactly this.\n\n" +
				"Examples:\n" +
				"- '(?:ab){2}' matches 'abab'\n" +
				"- '(?:log|err)' matches 'log' or 'err'\n" +
				"- '(?:re)?do' matches 'do' or 'redo'\n\n" +
				"This is particularly useful for:\n" +
				"- Grouping alternatives\n" +
				"- Applying quantifiers\n" +
				"- Performance optimization\n\n" +
				"Test text:\n" +
				"✗ ha (single pair)\n" +
				"✓ haha (two pairs)\n" +
				"✗ hahaha (three pairs)\n" +
				"✗ ahah (wrong order)",
			Task:        "Write a pattern using non-capturing group to match 'ha' repeated twice",
			TestCases: []models.TestCase{
				{Text: "ha", Expected: false},
				{Text: "haha", Expected: true},
				{Text: "hahaha", Expected: false},
				{Text: "ahahah", Expected: false},
			},
		},
		{
			Title:       "Escaping Special Characters",
			Description: "Let's learn about matching special regex characters literally!\n\n" +
				"Some characters have special meanings in regex. To match them literally, we need to escape them with a backslash (\\).\n\n" +
				"Special characters that need escaping:\n" +
				"- . * + ? [ ] ( ) { } ^ $ \\ |\n\n" +
				"Examples:\n" +
				"- '\\$5' matches a dollar sign and 5\n" +
				"- 'file\\.txt' matches 'file.txt'\n" +
				"- '2 \\+ 2' matches '2 + 2'\n\n" +
				"Test text:\n" +
				"✓ cat* (cat followed by asterisk)\n" +
				"✗ cat (no asterisk)\n" +
				"✗ cat? (wrong special character)\n" +
				"✗ cats (wrong character)",
			Task:        "Write a pattern that matches 'cat*' literally (including the asterisk)",
			TestCases: []models.TestCase{
				{Text: "cat*", Expected: true},
				{Text: "cat", Expected: false},
				{Text: "cattt", Expected: false},
				{Text: "cat?", Expected: false},
			},
		},
		{
			Title:       "Character Class Negation Shortcuts",
			Description: "Let's learn about shortcuts for matching what we don't want!\n\n" +
				"The uppercase versions of shortcuts match the opposite of their lowercase counterparts:\n" +
				"- \\D matches any non-digit\n" +
				"- \\W matches any non-word character\n" +
				"- \\S matches any non-whitespace\n\n" +
				"Examples:\n" +
				"- '\\D\\d' matches '@9', 'x4', but not '12'\n" +
				"- '\\W\\w' matches '#a', '!b', but not 'ab'\n" +
				"- '\\S\\s' matches 'a ', 'x\t', but not '  '\n\n" +
				"Test text:\n" +
				"✓ a1 (letter then digit)\n" +
				"✓ #5 (symbol then digit)\n" +
				"✗ 12 (digit then digit)\n" +
				"✗ ab (no digits)\n" +
				"✗ 1a (wrong order)",
			Task:        "Write a pattern that matches any single non-digit followed by any single digit",
			TestCases: []models.TestCase{
				{Text: "a1", Expected: true},
				{Text: "!2", Expected: true},
				{Text: ".5", Expected: true},
				{Text: "12", Expected: false},
				{Text: "aa", Expected: false},
				{Text: "1a", Expected: false},
			},
		},
		{
			Title:       "Greedy vs Lazy Quantifiers",
			Description: "Let's learn about different ways quantifiers can match!\n\n" +
				"By default, quantifiers (*, +, ?, {n,m}) are 'greedy' - they match as much as possible. " +
				"Adding a ? after a quantifier makes it 'lazy' - matching as little as possible.\n\n" +
				"Examples:\n" +
				"- '[0-9]*' greedy: '123' matches '123'\n" +
				"- '[0-9]*?' lazy: '123' matches '1'\n" +
				"- '<p>.*</p>' greedy: matches entire paragraphs\n" +
				"- '<p>.*?</p>' lazy: matches individual paragraphs\n\n" +
				"This is particularly useful for:\n" +
				"- Matching balanced delimiters\n" +
				"- Extracting specific content\n" +
				"- Controlled matching\n\n" +
				"Test text:\n" +
				"✓ <tag> (simple tag)\n" +
				"✓ <div> (another tag)\n" +
				"✗ <tag>content</tag> (too much content)\n" +
				"✗ tag (no brackets)",
			Task:        "Write a pattern that matches text between < and > brackets, taking the smallest possible match",
			TestCases: []models.TestCase{
				{Text: "<tag>", Expected: true},
				{Text: "<>", Expected: true},
				{Text: "tag", Expected: false},
				{Text: "<tag>value</tag>", Expected: true},
			},
		},
		{
			Title:       "Multiline Mode",
			Description: "Let's learn about handling multiple lines of text!\n\n" +
				"The (?m) flag changes how ^ and $ work - they match the start and end of each line instead of the whole text.\n\n" +
				"Examples:\n" +
				"- '(?m)^start' matches 'start' at beginning of any line\n" +
				"- '(?m)end$' matches 'end' at end of any line\n" +
				"- '(?m)^$' matches empty lines\n\n" +
				"This is particularly useful for:\n" +
				"- Processing text files\n" +
				"- Line-by-line validation\n" +
				"- Multi-line search\n\n" +
				"Test text:\n" +
				"✓ first line\n" +
				"✓ another line\n" +
				"✗ inline text\n" +
				"✗ text inline",
			Task:        "Write a pattern in multiline mode that matches 'line' at the end of any line",
			TestCases: []models.TestCase{
				{Text: "first line\n", Expected: true},
				{Text: "line\n", Expected: true},
				{Text: "line break", Expected: false},
				{Text: "inline\n", Expected: false},
			},
		},
		{
			Title:       "Case Insensitive Matching",
			Description: "Let's learn about ignoring letter case!\n\n" +
				"The (?i) flag makes your pattern match regardless of uppercase or lowercase letters.\n\n" +
				"Examples:\n" +
				"- '(?i)hi' matches 'hi', 'Hi', 'HI', 'hI'\n" +
				"- '(?i)yes' matches 'yes', 'YES', 'Yes', 'yEs'\n" +
				"- '(?i)ok' matches 'ok', 'OK', 'Ok', 'oK'\n\n" +
				"This is particularly useful for:\n" +
				"- User input matching\n" +
				"- Search functionality\n" +
				"- Text processing\n\n" +
				"Test text:\n" +
				"✓ CAT (uppercase)\n" +
				"✓ cat (lowercase)\n" +
				"✓ Cat (mixed case)\n" +
				"✗ dog (wrong word)",
			Task:        "Write a pattern that matches 'cat' regardless of letter case",
			TestCases: []models.TestCase{
				{Text: "CAT", Expected: true},
				{Text: "cat", Expected: true},
				{Text: "Cat", Expected: true},
				{Text: "cAt", Expected: true},
				{Text: "dog", Expected: false},
			},
		},
		{
			Title:       "Unicode Categories",
			Description: "Let's learn about matching characters from any language!\n\n" +
				"Unicode categories help you match broad types of characters:\n" +
				"- \\p{L} matches letters from any language\n" +
				"- \\p{N} matches numbers from any numbering system\n" +
				"- \\p{P} matches punctuation marks\n\n" +
				"Examples:\n" +
				"- '\\p{L}+' matches words in any language\n" +
				"- '\\p{N}+' matches numbers in any script\n" +
				"- '\\p{P}+' matches punctuation sequences\n\n" +
				"Test text:\n" +
				"✓ A1 (Latin letter + number)\n" +
				"✓ Б2 (Cyrillic letter + number)\n" +
				"✓ 漢3 (Chinese character + number)\n" +
				"✗ 1A (number + letter)\n" +
				"✗ AA (no number)\n" +
				"✗ 12 (no letter)",
			Task:        "Write a pattern that matches any letter from any language followed by a number",
			TestCases: []models.TestCase{
				{Text: "A1", Expected: true},
				{Text: "Б2", Expected: true},
				{Text: "漢3", Expected: true},
				{Text: "1A", Expected: false},
				{Text: "AA", Expected: false},
				{Text: "11", Expected: false},
				{Text: "🎯4", Expected: false},
			},
		},
		{
			Title:       "Backreferences",
			Description: "Let's learn about referring back to matched content!\n\n" +
				"When you capture text in parentheses (), you can refer back to it later in your pattern using \\1, \\2, etc.\n\n" +
				"Examples:\n" +
				"- '(hi)-\\1' matches 'hi-hi'\n" +
				"- '(\\w)\\1' matches 'aa', 'bb', etc.\n" +
				"- '(\\d\\d)=\\1' matches '42=42'\n\n" +
				"This is particularly useful for:\n" +
				"- Finding repeated content\n" +
				"- Matching paired items\n" +
				"- Pattern validation\n\n" +
				"Test text:\n" +
				"✓ aa (letter repeated)\n" +
				"✓ bb (different letter repeated)\n" +
				"✗ ab (different letters)\n" +
				"✗ a (single letter)",
			Task:        "Write a pattern that matches any letter followed by the same letter",
			TestCases: []models.TestCase{
				{Text: "aa", Expected: true},
				{Text: "bb", Expected: true},
				{Text: "cc", Expected: true},
				{Text: "ab", Expected: false},
				{Text: "a", Expected: false},
			},
		},
		{
			Title:       "Named Groups",
			Description: "Let's learn about giving names to captured groups!\n\n" +
				"Named groups (?P<name>pattern) let you give meaningful names to parts of your pattern. " +
				"You can refer back to them with \\k<name>.\n\n" +
				"Examples:\n" +
				"- '(?P<year>\\d{4})-\\k<year>' matches '2024-2024'\n" +
				"- '(?P<tag>\\w+)</\\k<tag>>' matches 'div</div>'\n" +
				"- '(?P<char>.)\\k<char>\\k<char>' matches 'aaa'\n\n" +
				"This is particularly useful for:\n" +
				"- Self-documenting patterns\n" +
				"- Complex backreferences\n" +
				"- Pattern maintenance\n\n" +
				"Test text:\n" +
				"✓ cat=cat (same word both sides)\n" +
				"✓ dog=dog (same word both sides)\n" +
				"✗ cat=dog (different words)\n" +
				"✗ cat=cats (different forms)",
			Task:        "Write a pattern with a named group 'word' that matches the same word before and after an equals sign",
			TestCases: []models.TestCase{
				{Text: "cat=cat", Expected: true},
				{Text: "dog=dog", Expected: true},
				{Text: "cat=dog", Expected: false},
				{Text: "dog=cat", Expected: false},
			},
		},
		{
			Title:       "Whitespace Patterns",
			Description: "Let's learn about matching different types of whitespace!\n\n" +
				"Regex provides several ways to match whitespace:\n" +
				"- \\t matches a tab character\n" +
				"- \\n matches a newline\n" +
				"- \\r matches a carriage return\n" +
				"- \\s matches any whitespace\n\n" +
				"Examples:\n" +
				"- 'a\\tb' matches 'a' and 'b' separated by a tab\n" +
				"- 'end\\n' matches 'end' followed by newline\n" +
				"- 'a\\s+b' matches 'a' and 'b' with any whitespace between\n\n" +
				"Test text:\n" +
				"✓ word\tword (tab between)\n" +
				"✗ word word (space between)\n" +
				"✗ word\nword (newline between)\n" +
				"✗ wordword (no separation)",
			Task:        "Write a pattern that matches 'word' followed by a tab followed by 'word'",
			TestCases: []models.TestCase{
				{Text: "word\tword", Expected: true},
				{Text: "word word", Expected: false},
				{Text: "word\nword", Expected: false},
				{Text: "word", Expected: false},
			},
		},
	}
}

func GetPracticeProblems() []models.PracticeProblem {
	return []models.PracticeProblem{
		{
			Title:       "IP Address",
			Description: "Write a pattern to match IPv4 addresses.\nEach number should be between 0-255.",
			Examples:    "Valid: 192.168.1.1, 10.0.0.0\nInvalid: 256.1.2.3, 1.2.3.4.5",
			TestCases: []models.TestCase{
				{Text: "192.168.1.1", Expected: true},
				{Text: "10.0.0.0", Expected: true},
				{Text: "256.1.2.3", Expected: false},
				{Text: "1.2.3.4.5", Expected: false},
			},
		},
		{
			Title:       "HTML Color Codes",
			Description: "Write a pattern to match HTML hex color codes.",
			Examples:    "Valid: #FFF, #123456\nInvalid: #XYZ, #12345",
			TestCases: []models.TestCase{
				{Text: "#FFF", Expected: true},
				{Text: "#123456", Expected: true},
				{Text: "#XYZ", Expected: false},
				{Text: "#12345", Expected: false},
			},
		},
		{
			Title:       "Time Format",
			Description: "Write a pattern to match 24-hour time format (HH:MM).",
			Examples:    "Valid: 13:45, 09:30\nInvalid: 24:00, 12:60",
			TestCases: []models.TestCase{
				{Text: "13:45", Expected: true},
				{Text: "09:30", Expected: true},
				{Text: "24:00", Expected: false},
				{Text: "12:60", Expected: false},
			},
		},
		{
			Title:       "Variable Names",
			Description: "Write a pattern to validate JavaScript variable names.\nMust start with letter/$/_, followed by letters/numbers/$/_",
			Examples:    "Valid: myVar, $price, _hidden\nInvalid: 123var, my-var, class@",
			TestCases: []models.TestCase{
				{Text: "myVar", Expected: true},
				{Text: "$price", Expected: true},
				{Text: "_hidden", Expected: true},
				{Text: "123var", Expected: false},
				{Text: "my-var", Expected: false},
				{Text: "class@", Expected: false},
			},
		},
		{
			Title:       "Version Numbers",
			Description: "Write a pattern to match semantic version numbers (x.y.z format).\nEach number can have 1-3 digits.",
			Examples:    "Valid: 1.0.0, 2.10.5, 10.20.30\nInvalid: 1.0, 1.0.0.0, 01.02.03",
			TestCases: []models.TestCase{
				{Text: "1.0.0", Expected: true},
				{Text: "2.10.5", Expected: true},
				{Text: "10.20.30", Expected: true},
				{Text: "1.0", Expected: false},
				{Text: "1.0.0.0", Expected: false},
				{Text: "01.02.03", Expected: false},
			},
		},
		{
			Title:       "Log Level Extraction",
			Description: "Write a pattern to match log levels in brackets at start of line.\nShould extract: [ERROR], [INFO], [WARN], [DEBUG]",
			Examples:    "Valid: [ERROR] Failed to connect\nInvalid: [FATAL] Error occurred",
			TestCases: []models.TestCase{
				{Text: "[ERROR] Failed to connect", Expected: true},
				{Text: "[INFO] Server started", Expected: true},
				{Text: "[WARN] Disk space low", Expected: true},
				{Text: "[DEBUG] Processing request", Expected: true},
				{Text: "[FATAL] Error occurred", Expected: false},
				{Text: "ERROR: System crash", Expected: false},
			},
		},
		{
			Title:       "File Extensions",
			Description: "Write a pattern to match common web file extensions at end of string.\nMatch: .jpg, .jpeg, .png, .gif, .webp",
			Examples:    "Valid: photo.jpg, image.png\nInvalid: document.pdf, script.js",
			TestCases: []models.TestCase{
				{Text: "photo.jpg", Expected: true},
				{Text: "image.jpeg", Expected: true},
				{Text: "icon.png", Expected: true},
				{Text: "animation.gif", Expected: true},
				{Text: "photo.webp", Expected: true},
				{Text: "document.pdf", Expected: false},
				{Text: "script.js", Expected: false},
			},
		},
		{
			Title:       "URL Path Parameters",
			Description: "Write a pattern to match URL path parameters.\nFormat: /:paramName/",
			Examples:    "Valid: /users/:id/profile, /:category/:productId\nInvalid: /users/123, /:1id/",
			TestCases: []models.TestCase{
				{Text: "/users/:id/profile", Expected: true},
				{Text: "/:category/:productId", Expected: true},
				{Text: "/users/123", Expected: false},
				{Text: "/:1id/", Expected: false},
			},
		},
		{
			Title:       "CSS Color Values",
			Description: "Write a pattern to match CSS RGB color values.\nFormat: rgb(0-255, 0-255, 0-255)",
			Examples:    "Valid: rgb(255, 128, 0), rgb(0, 0, 0)\nInvalid: rgb(300, 0, 0), rgb(0,0,0)",
			TestCases: []models.TestCase{
				{Text: "rgb(255, 128, 0)", Expected: true},
				{Text: "rgb(0, 0, 0)", Expected: true},
				{Text: "rgb(300, 0, 0)", Expected: false},
				{Text: "rgb(0,0,0)", Expected: false},
			},
		},
		{
			Title:       "JSON Property",
			Description: "Write a pattern to match JSON property names and values.\nFormat: \"propertyName\": value",
			Examples:    "Valid: \"name\": \"John\", \"age\": 30\nInvalid: name: \"John\", \"age\":30",
			TestCases: []models.TestCase{
				{Text: "\"name\": \"John\"", Expected: true},
				{Text: "\"age\": 30", Expected: true},
				{Text: "name: \"John\"", Expected: false},
				{Text: "\"age\":30", Expected: false},
			},
		},
		{
			Title:       "Git Commit Hash",
			Description: "Write a pattern to match Git commit hashes.\n40 characters of hexadecimal (short form: 7+ chars)",
			Examples:    "Valid: a1b2c3d4e5f6..., 1234567\nInvalid: xyz123, 123456",
			TestCases: []models.TestCase{
				{Text: "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0", Expected: true},
				{Text: "1234567", Expected: true},
				{Text: "xyz123", Expected: false},
				{Text: "123456", Expected: false},
			},
		},
		{
			Title:       "Docker Image Tags",
			Description: "Write a pattern to match Docker image tags.\nFormat: repository/image:tag",
			Examples:    "Valid: ubuntu:20.04, nginx:latest\nInvalid: ubuntu:, :latest",
			TestCases: []models.TestCase{
				{Text: "ubuntu:20.04", Expected: true},
				{Text: "nginx:latest", Expected: true},
				{Text: "mysql:5.7.33", Expected: true},
				{Text: "ubuntu:", Expected: false},
				{Text: ":latest", Expected: false},
			},
		},
		{
			Title:       "Function Parameters",
			Description: "Write a pattern to match function parameters in JavaScript.\nMatch parameters between parentheses.",
			Examples:    "Valid: (x, y), (name = 'default')\nInvalid: (), (x,y)",
			TestCases: []models.TestCase{
				{Text: "(x, y)", Expected: true},
				{Text: "(name = 'default')", Expected: true},
				{Text: "()", Expected: false},
				{Text: "(x,y)", Expected: false},
			},
		},
		{
			Title:       "Database Connection String",
			Description: "Write a pattern to match PostgreSQL connection strings.\nFormat: postgresql://user:pass@host:port/dbname",
			Examples:    "Valid: postgresql://user:pass@localhost:5432/mydb\nInvalid: mysql://user:pass@localhost/db",
			TestCases: []models.TestCase{
				{Text: "postgresql://user:pass@localhost:5432/mydb", Expected: true},
				{Text: "postgresql://admin:secret@db.host:5432/prod", Expected: true},
				{Text: "mysql://user:pass@localhost/db", Expected: false},
				{Text: "postgresql://user@localhost/db", Expected: false},
			},
		},
		{
			Title:       "API Endpoints",
			Description: "Write a pattern to match RESTful API endpoints.\nFormat: /api/v1/resource/:id",
			Examples:    "Valid: /api/v1/users/123, /api/v2/products\nInvalid: /api/v0/users, /api/users",
			TestCases: []models.TestCase{
				{Text: "/api/v1/users/123", Expected: true},
				{Text: "/api/v2/products", Expected: true},
				{Text: "/api/v0/users", Expected: false},
				{Text: "/api/users", Expected: false},
			},
		},
		{
			Title:       "Environment Variables",
			Description: "Write a pattern to match environment variables in shell scripts.\nFormat: $VARIABLE or ${VARIABLE}",
			Examples:    "Valid: $HOME, ${PATH}\nInvalid: $123, ${1VAR}",
			TestCases: []models.TestCase{
				{Text: "$HOME", Expected: true},
				{Text: "${PATH}", Expected: true},
				{Text: "$123", Expected: false},
				{Text: "${1VAR}", Expected: false},
			},
		},
		{
			Title:       "HTML Data Attributes",
			Description: "Write a pattern to match HTML5 data attributes.\nFormat: data-* attributes",
			Examples:    "Valid: data-user-id, data-color\nInvalid: data-, data@test",
			TestCases: []models.TestCase{
				{Text: "data-user-id", Expected: true},
				{Text: "data-color", Expected: true},
				{Text: "data-", Expected: false},
				{Text: "data@test", Expected: false},
			},
		},
		{
			Title:       "CSS Media Queries",
			Description: "Write a pattern to match CSS media query breakpoints.\nFormat: @media (min-width: XXXpx)",
			Examples:    "Valid: @media (min-width: 768px)\nInvalid: @media screen, @media (width: 768px)",
			TestCases: []models.TestCase{
				{Text: "@media (min-width: 768px)", Expected: true},
				{Text: "@media (min-width: 1200px)", Expected: true},
				{Text: "@media screen", Expected: false},
				{Text: "@media (width: 768px)", Expected: false},
			},
		},
		{
			Title:       "Package Version Range",
			Description: "Write a pattern to match npm package version ranges.\nFormat: ^1.2.3, ~1.2.3, >=1.2.3",
			Examples:    "Valid: ^1.0.0, ~2.1.0, >=3.0.0\nInvalid: >>1.0.0, ^1.0",
			TestCases: []models.TestCase{
				{Text: "^1.0.0", Expected: true},
				{Text: "~2.1.0", Expected: true},
				{Text: ">=3.0.0", Expected: true},
				{Text: ">>1.0.0", Expected: false},
				{Text: "^1.0", Expected: false},
			},
		},
		{
			Title:       "Base64 Strings",
			Description: "Write a pattern to match Base64 encoded strings.\nMust be multiple of 4 chars, using A-Z, a-z, 0-9, +, /, =",
			Examples:    "Valid: SGVsbG8=, YWJjZA==\nInvalid: abc=, Hello!",
			TestCases: []models.TestCase{
				{Text: "SGVsbG8=", Expected: true},
				{Text: "YWJjZA==", Expected: true},
				{Text: "abc=", Expected: false},
				{Text: "Hello!", Expected: false},
			},
		},
		{
			Title:       "JWT Token",
			Description: "Write a pattern to match JWT tokens.\nThree Base64 sections separated by dots",
			Examples:    "Valid: eyJhbGci.eyJzdWIi.TJVA95Or\nInvalid: abc.def, a.b.c.d",
			TestCases: []models.TestCase{
				{Text: "eyJhbGci.eyJzdWIi.TJVA95Or", Expected: true},
				{Text: "header.payload.signature", Expected: true},
				{Text: "abc.def", Expected: false},
				{Text: "a.b.c.d", Expected: false},
			},
		},
		{
			Title:       "MongoDB ObjectId",
			Description: "Write a pattern to match MongoDB ObjectId.\n24 character hex string",
			Examples:    "Valid: 507f1f77bcf86cd799439011\nInvalid: 507f1f77bcf86cd79943901, 507f1f77bcf86cd7994390111",
			TestCases: []models.TestCase{
				{Text: "507f1f77bcf86cd799439011", Expected: true},
				{Text: "abcdef0123456789abcdef01", Expected: true},
				{Text: "507f1f77bcf86cd79943901", Expected: false},
				{Text: "507f1f77bcf86cd7994390111", Expected: false},
			},
		},
		{
			Title:       "Kubernetes Resource Names",
			Description: "Write a pattern to match valid Kubernetes resource names.\nLowercase, numbers, -, up to 253 chars",
			Examples:    "Valid: my-app-pod, nginx-123\nInvalid: My_Pod, pod@123",
			TestCases: []models.TestCase{
				{Text: "my-app-pod", Expected: true},
				{Text: "nginx-123", Expected: true},
				{Text: "My_Pod", Expected: false},
				{Text: "pod@123", Expected: false},
			},
		},
		{
			Title:       "GraphQL Fields",
			Description: "Write a pattern to match GraphQL field selections.\nFormat: fieldName or fieldName { subfield }",
			Examples:    "Valid: userName, user { id name }\nInvalid: 123field, user{id}",
			TestCases: []models.TestCase{
				{Text: "userName", Expected: true},
				{Text: "user { id name }", Expected: true},
				{Text: "123field", Expected: false},
				{Text: "user{id}", Expected: false},
			},
		},
		{
			Title:       "CI/CD Variables",
			Description: "Write a pattern to match CI/CD pipeline variables.\nFormat: ${VAR_NAME} or $VAR_NAME",
			Examples:    "Valid: ${BUILD_ID}, $DEPLOY_ENV\nInvalid: $123, ${1BUILD}",
			TestCases: []models.TestCase{
				{Text: "${BUILD_ID}", Expected: true},
				{Text: "$DEPLOY_ENV", Expected: true},
				{Text: "$123", Expected: false},
				{Text: "${1BUILD}", Expected: false},
			},
		},
	}
}
