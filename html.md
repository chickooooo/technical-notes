## Index
- [DOCTYPE Declaration](#doctype-declaration)
- [HTML Structure](#html-structure)
- [Body Tag](#body-tag)
- [Heading Tag](#heading-tag)
- [Paragraph Tag](#paragraph-tag)
- [Anchor (Link) Tag](#anchor-link-tag)
- [Image Tag](#image-tag)
- [Line Break Tag](#line-break-tag)
- [Horizontal Rule Tag](#horizontal-rule-tag)
- [Preformatted Text Tag](#preformatted-text-tag)
- [Style Attribute](#style-attribute)
- [Bold Text - b Tag](#bold-text---b-tag)
- [Strong Text - strong Tag](#strong-text---strong-tag)
- [Italic Text - i Tag](#italic-text---i-tag)
- [Emphasized Text - em Tag](#emphasized-text---em-tag)
- [Small Text - small Tag](#small-text---small-tag)
- [Mark Tag](#mark-tag)
- [Deleted Text - del Tag](#deleted-text---del-tag)
- [Inserted Text - ins Tag](#inserted-text---ins-tag)
- [Subscript Text - sub Tag](#subscript-text---sub-tag)
- [Superscript Text - sup Tag](#superscript-text---sup-tag)
- [HTML Comments](#html-comments)
- [Target Attribute](#target-attribute)
- [Favicon](#favicon)
- [HTML Tables](#html-tables)
- [Unordered and Ordered Lists](#unordered-and-ordered-lists)
- [Description List](#description-list)
- [Block-level Elements](#block-level-elements)
- [Inline Elements](#inline-elements)
- [Head Element](#head-element)
- [Title Element](#title-element)
- [Style Element](#style-element)
- [Link Element](#link-element)
- [Meta Element](#meta-element)
- [Header Element](#header-element)
- [Nav Element](#nav-element)
- [Section Element](#section-element)
- [Article Element](#article-element)
- [Aside Element](#aside-element)
- [Footer Element](#footer-element)
- [Details and Summary](#details-and-summary)
- [Figure and Figcaption](#figure-and-figcaption)
- [Main Element](#main-element)
- [Form Element](#form-element)
- [Input Element](#input-element)
- [Label Element](#label-element)
- [Script Tag](#script-tag)
- [Noscript Tag](#noscript-tag)
- [Iframe Tag](#iframe-tag)
- [Canvas Tag](#canvas-tag)
- [SVG Tag](#svg-tag)
- [Audio Tag](#audio-tag)
- [Video Tag](#video-tag)
- [Time Tag](#time-tag)
- [Datalist Tag](#datalist-tag)

<br>
<br>
<br>

### DOCTYPE Declaration
- Declares the document type and version of HTML.
- Helps browsers render the page correctly.

<br>

```html
<!DOCTYPE html>
````

<br>
<br>
<br>

### HTML Structure

* The entire HTML document starts with `<html>` and ends with `</html>`.

<br>

```html
<html>
  <!-- Content goes here -->
</html>
```

<br>
<br>
<br>

### Body Tag

* Contains the visible content of the web page.
* Placed between `<body>` and `</body>`.

<br>

```html
<body>
  <!-- Visible content -->
</body>
```

<br>
<br>
<br>

### Heading Tag

* Defines headings from `<h1>` (most important) to `<h6>` (least important).

<br>

```html
<h1>This is a primary heading</h1>
<h2>This is a secondary heading</h2>
```

<br>
<br>
<br>

### Paragraph Tag

* Used to define blocks of text or paragraphs.

<br>

```html
<p>This is a paragraph.</p>
```

<br>
<br>
<br>

### Anchor (Link) Tag

* Defines hyperlinks using the `<a>` tag.
* The `href` attribute specifies the link's destination.

<br>

```html
<a href="https://www.w3schools.com">This is a link</a>
```

<br>
<br>
<br>

### Image Tag
- Embeds an image in a web page using the `<img>` tag.
- Empty tag with required `src` and `alt` attributes.

<br>

```html
<img src="w3schools.jpg" alt="W3Schools.com" width="104" height="142">
````

<br>
<br>
<br>

### Line Break Tag

* Adds a line break.
* Empty element with no closing tag.

<br>

```html
<p>This is a sentence.<br>This is a new line.</p>
```

<br>
<br>
<br>

### Horizontal Rule Tag

* Creates a horizontal line to separate content.
* Used for thematic breaks.

<br>

```html
<hr>
```

<br>
<br>
<br>

### Preformatted Text Tag

* Displays preformatted text with fixed-width font.
* Preserves spaces and line breaks.

<br>

```html
<pre>
  Line 1
    Line 2
      Line 3
</pre>
```

<br>
<br>
<br>

### Style Attribute

* Adds inline styles to HTML elements.
* Uses CSS property-value pairs.

<br>

```html
<p style="color:blue; font-size:18px;">Styled paragraph</p>
```

<br>
<br>
<br>

### Bold Text - b Tag

* Makes text bold without implying importance.

<br>

```html
<b>This text is bold</b>
```

<br>
<br>
<br>

### Strong Text - strong Tag

* Indicates strong importance.
* Also renders as bold, but with semantic emphasis.

<br>

```html
<strong>This text is important</strong>
```

<br>
<br>
<br>

### Italic Text - i Tag

* Defines text in an alternate voice or mood.
* Typically displayed in italics.

<br>

```html
<i>This is italic text</i>
```

<br>
<br>
<br>

### Emphasized Text - em Tag

* Emphasizes text, usually shown in italics.
* Screen readers read it with stress.

<br>

```html
<em>This text is emphasized</em>
```

<br>
<br>
<br>

### Small Text - small Tag

* Displays text in a smaller size.

<br>

```html
<small>This is some smaller text.</small>
```

<br>
<br>
<br>

### Mark Tag
- Highlights or marks text for reference.

<br>

```html
<p>Do not forget to buy <mark>milk</mark> today.</p>
````

<br>
<br>
<br>

### Deleted Text - del Tag

* Represents text that has been removed.
* Usually displayed with a strikethrough.

<br>

```html
<p>My favorite color is <del>blue</del> red.</p>
```

<br>
<br>
<br>

### Inserted Text - ins Tag

* Represents newly added text.
* Usually displayed with an underline.

<br>

```html
<p>My favorite color is <del>blue</del> <ins>red</ins>.</p>
```

<br>
<br>
<br>

### Subscript Text - sub Tag

* Displays text slightly below the normal line.
* Commonly used for chemical formulas.

<br>

```html
<p>This is <sub>subscripted</sub> text.</p>
```

<br>
<br>
<br>

### Superscript Text - sup Tag

* Displays text slightly above the normal line.
* Commonly used for footnotes.

<br>

```html
<p>This is <sup>superscripted</sup> text.</p>
```

<br>
<br>
<br>

### HTML Comments

* Comments are not displayed in the browser.
* Useful for notes or hiding code.

<br>

```html
<!-- Write your comments here -->

<!--
<p>Look at this cool image:</p>
<img border="0" src="pic_trulli.jpg" alt="Trulli">
-->
```

<br>
<br>
<br>

### Target Attribute

* Specifies where to open the linked document.
* Common values:

  * `_self`: opens in the same tab (default)
  * `_blank`: opens in a new tab/window

<br>

```html
<a href="https://www.w3schools.com/" target="_blank">Visit W3Schools!</a>
```

<br>
<br>
<br>

### Favicon

* Adds a small icon in the browser tab.
* The `rel="icon"` link tag is placed in the `<head>` section.

<br>

```html
<link rel="icon" type="image/x-icon" href="/images/favicon.ico">
```

<br>
<br>
<br>

### HTML Tables
- `<table>` creates a table.
- `<tr>` defines a row, `<td>` a data cell, `<th>` a header cell.
- `<thead>`, `<tbody>`, `<tfoot>` group parts of the table.
- `<caption>` adds a title to the table.

<br>

```html
<table>
  <caption>Monthly Sales</caption>
  <thead>
    <tr><th>Month</th><th>Sales</th></tr>
  </thead>
  <tbody>
    <tr><td>January</td><td>$1000</td></tr>
  </tbody>
</table>
````

<br>
<br>
<br>

### Unordered and Ordered Lists

* `<ul>` creates an unordered (bulleted) list.
* `<ol>` creates an ordered (numbered) list.
* `<li>` defines each list item.

<br>

```html
<ul>
  <li>Coffee</li>
  <li>Tea</li>
</ul>

<ol>
  <li>First</li>
  <li>Second</li>
</ol>
```

<br>
<br>
<br>

### Description List

* `<dl>` creates a description list.
* `<dt>` defines the term.
* `<dd>` describes the term.

<br>

```html
<dl>
  <dt>Coffee</dt>
  <dd>- black hot drink</dd>
  <dt>Milk</dt>
  <dd>- white cold drink</dd>
</dl>
```

<br>
<br>
<br>

### Block-level Elements

* Start on a new line and take up full width.
* Common examples: `<div>`, `<p>`.

<br>

```html
<div>This is a block-level element</div>
```

<br>
<br>
<br>

### Inline Elements

* Do not start on a new line.
* Take up only as much space as needed.
* Common example: `<span>`.

<br>

```html
<span style="color:blue;">Blue text</span>
```

<br>
<br>
<br>

### Head Element

* Contains metadata for the document.
* Placed inside `<html>`, before `<body>`.

<br>

```html
<head>
  <!-- Metadata and links go here -->
</head>
```

<br>
<br>
<br>

### Title Element

* Defines the page title shown in the browser tab.
* Must be text-only.

<br>

```html
<head>
  <title>HTML Tutorial</title>
</head>
```

<br>
<br>
<br>

### Style Element

* Adds internal CSS to a page.

<br>

```html
<style>
  body {background-color: powderblue;}
  h1 {color: red;}
  p {color: blue;}
</style>
```

<br>
<br>
<br>

### Link Element

* Links external resources like stylesheets.

<br>

```html
<link rel="stylesheet" href="mystyle.css">
```

<br>
<br>
<br>

### Meta Element

* Provides metadata like charset, description, and viewport.

<br>

```html
<meta charset="UTF-8">
<meta name="description" content="Free Web tutorials">
```

<br>
<br>
<br>

### Header Element

* Defines introductory content or navigation.

<br>

```html
<header>
  <h1>Welcome</h1>
  <nav>...</nav>
</header>
```

<br>
<br>
<br>

### Nav Element

* Contains navigation links.

<br>

```html
<nav>
  <a href="#home">Home</a>
  <a href="#about">About</a>
</nav>
```

<br>
<br>
<br>

### Section Element

* Groups related content under a common theme.

<br>

```html
<section>
  <h2>Blog Posts</h2>
  <p>Post content...</p>
</section>
```

<br>
<br>
<br>

### Article Element

* Contains independent, self-contained content.

<br>

```html
<article>
  <h2>News Headline</h2>
  <p>Full story here...</p>
</article>
```

<br>
<br>
<br>

### Aside Element

* Defines content aside from the main content (like a sidebar).

<br>

```html
<aside>
  <h4>Related Links</h4>
</aside>
```

<br>
<br>
<br>

### Footer Element

* Defines the footer for a page or section.

<br>

```html
<footer>
  <p>© 2025 My Website</p>
</footer>
```

<br>
<br>
<br>

### Details and Summary

* `<details>` allows toggling hidden details.
* `<summary>` is the visible heading.

<br>

```html
<details>
  <summary>More Info</summary>
  <p>This is hidden until expanded.</p>
</details>
```

<br>
<br>
<br>

### Figure and Figcaption

* `<figure>` wraps media and related content.
* `<figcaption>` provides a caption.

<br>

```html
<figure>
  <img src="pic_trulli.jpg" alt="Trulli">
  <figcaption>Fig1. - Trulli, Puglia, Italy.</figcaption>
</figure>
```

<br>
<br>
<br>

### Main Element

* Specifies the primary content of the document.

<br>

```html
<main>
  <h1>Main Content</h1>
</main>
```

<br>
<br>
<br>

### Form Element

* Wraps user input elements for form submission.

<br>

```html
<form action="/submit">
  <!-- Input elements go here -->
</form>
```

<br>
<br>
<br>

### Input Element

* Creates various types of form inputs.
* `name` attribute is required for submission.

<br>

```html
<input type="text" name="username">
<input type="radio" name="gender" value="male"> Male
<input type="checkbox" name="vehicle" value="bike"> Bike
<input type="submit" value="Submit">
<input type="button" value="Click Me">
```

<br>
<br>
<br>

### Label Element

* Defines a label for an input element.
* Improves form accessibility.

<br>

```html
<label for="name">Name:</label>
<input type="text" id="name" name="name">
```

<br>
<br>
<br>

### Script Tag
- Embeds or links JavaScript in an HTML page.
- Can be placed in `<head>` or `<body>`.

<br>

```html
<script>
  alert("Hello, world!");
</script>

<script src="script.js"></script>
````

<br>
<br>
<br>

### Noscript Tag

* Provides fallback content for users with JavaScript disabled.

<br>

```html
<noscript>
  Your browser does not support JavaScript or it is disabled.
</noscript>
```

<br>
<br>
<br>

### Iframe Tag

* Embeds another HTML page within the current page.

<br>

```html
<iframe src="https://www.example.com" width="300" height="200"></iframe>
```

<br>
<br>
<br>

### Canvas Tag

* Provides a space for drawing graphics via JavaScript.
* Requires scripting to render visuals.

<br>

```html
<canvas id="myCanvas" width="200" height="100" style="border:1px solid #000;"></canvas>
```

<br>
<br>
<br>

### SVG Tag

* Defines vector-based graphics that scale cleanly.
* Can include shapes like circles, lines, and paths.

<br>

```html
<svg width="100" height="100">
  <circle cx="50" cy="50" r="40" stroke="green" fill="yellow" />
</svg>
```

<br>
<br>
<br>

### Audio Tag

* Embeds audio files.
* Supports multiple formats and playback controls.

<br>

```html
<audio controls>
  <source src="audio.mp3" type="audio/mpeg">
  Your browser does not support the audio element.
</audio>
```

<br>
<br>
<br>

### Video Tag

* Embeds video content.
* Supports controls, poster images, and multiple formats.

<br>

```html
<video width="320" height="240" controls>
  <source src="movie.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>
```

<br>
<br>
<br>

### Time Tag

* Represents time or a duration.
* Can include a `datetime` attribute for machine-readable format.

<br>

```html
<time datetime="2025-08-12">August 12, 2025</time>
```

<br>
<br>
<br>

### Datalist Tag

* Provides predefined options for an `<input>` element.
* Enables autocomplete functionality.

<br>

```html
<input list="browsers" name="browser">
<datalist id="browsers">
  <option value="Chrome">
  <option value="Firefox">
  <option value="Safari">
</datalist>
```
