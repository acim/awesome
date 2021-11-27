# CSS

- [Material Design Components](https://material.io/develop/web/components/animation/)
- [Creating Non-Rectangular Headers](https://css-tricks.com/creating-non-rectangular-headers/)
- [CSS clip-path maker](https://bennettfeely.com/clippy/)
- [List Animations](https://codepen.io/designcourse/pen/YzPmJxo)
- [Staggered Animations with CSS Custom Properties](https://cloudfour.com/thinks/staggered-animations-with-css-custom-properties/)
- [CSS Tricks](https://css-tricks.com/)
- [CSS Reference](http://tympanus.net/codrops/css_reference)
- [Can I use? - Check if browser supports some feature](https://caniuse.com/)
- [sanitize.css - CSS library that provides consistent, cross-browser default styling of HTML elements alongside useful defaults](https://github.com/csstools/sanitize.css)
- [Even more CSS secrets](https://www.youtube.com/watch?v=vs34f9FiHps)
- [Using @font-face](https://css-tricks.com/snippets/css/using-font-face/)
- [Hot CSS tips](https://equinusocio.dev/blog/hot-tips-css/)
- [Content CSS property - none, normal, string, image, counter, attr, open-quote](https://developer.mozilla.org/en-US/docs/Web/CSS/content)
- [:target pseudo selector](https://css-tricks.com/on-target/)
- [transform-origin property](https://css-tricks.com/almanac/properties/t/transform-origin/)
- [object-fit property - style images](https://css-tricks.com/almanac/properties/o/object-fit/)
- [Seamless Responsive Photo Grid](https://css-tricks.com/seamless-responsive-photo-grid/)
- [flex properties](https://css-tricks.com/almanac/properties/f/flex/)
- [Jonas' Resources for Building Beautiful Websites with HTML5, CSS3 and JavaScript](http://codingheroes.io/resources/)
- [modern-normalize](https://github.com/sindresorhus/modern-normalize)
- [Blueprint CSS - A lightweight layout library for building great responsive mobile first UIs that work everywhere](https://blueprintcss.dev/)
- [Responsive Grid System](http://www.responsivegridsystem.com/)
- [Fancy CSS Button Hover Animations](https://www.youtube.com/watch?v=cH0TC9gWiAg)
- [Fancy CSS Button Hover Animations](https://github.com/WebDevSimplified/Fancy-CSS-Button-Hover-Animations)
- [Social Media Accordion With CSS3 Transitions](https://www.youtube.com/watch?v=4M6qPoFWIxI)
- [Animista - Create CSS Animations](https://animista.net/)
- [Pattern.css - CSS Patterns](https://bansal.io/pattern-css)
- [98.css - Windows 98 Style UI](https://jdan.github.io/98.css/)
- [Guide to Responsive-Friendly CSS Columns](https://css-tricks.com/guide-responsive-friendly-css-columns/)
- [What CSS to prefix?](http://shouldiprefix.com/)
- [all property - resets all of the selected element’s properties](https://css-tricks.com/almanac/properties/a/all/)
- [Normalize CSS or CSS Reset?! - code to reset most important elements](https://medium.com/@elad/normalize-css-or-css-reset-9d75175c5d1e)
- [perspective property - gives an element a 3D-space by affecting the distance between the Z plane and the user](https://css-tricks.com/almanac/properties/p/perspective/)
- [backface-visibility property - rotate an element so what we think of as the front of an element no longer faces the screen](https://css-tricks.com/almanac/properties/b/backface-visibility/)
- [Cascading image using position sticky](https://codepen.io/drnaia100/full/LvbxNj)
- [:target pseudo selector - matches when the hash in the URL and the id of an element are the same](https://css-tricks.com/on-target/)
- [100% Pure HTML/CSS Page Navigation - :target](https://dev.to/hakash/100-pure-htmlcss-page-navigation---no-javascript-required-2em5)
- [How to Create Image Slider Using HTML and CSS](https://www.youtube.com/watch?v=FZrHoAUkHpE)
- [CSS Card Tricks - Cool transition cards and CSS-Tricks like avatars](https://www.youtube.com/watch?v=29deL9MFfWc)
- [CSS Frosted glass credit card](https://dev.to/dailydevtips1/css-frosted-glass-credit-card-3lak)
- [:is() pseudo class](https://css-tricks.com/almanac/selectors/i/is/)
- [display: flow-root - replacement for clearfix](https://css-tricks.com/display-flow-root/)
- [shape-outside - wrap text in circle, ellipse or other shape](https://css-tricks.com/almanac/properties/s/shape-outside/)
- [1-Line Layouts](https://1linelayouts.glitch.me/)
- [My Custom CSS Reset](https://www.joshwcomeau.com/css/custom-css-reset/)

## Gradients

- [Fun ways to animate CSS gradients](https://www.youtube.com/watch?v=f3mwKLXpOLk)

## Patterns

- [css-doodle - Web component for drawing patterns with CSS](https://css-doodle.com/)

## Tips & Tricks

### CSS Variables

```css
:root {
  --primary: #0f0;
  --secondary: #fff;
}
html {
  font-family: Arial, Helvetica, sans-serif;
  color: var(--primary);
}
```

### [Apply a natural box layout model to all elements, but allowing components to change](https://www.paulirish.com/2012/box-sizing-border-box-ftw/)

```css
html {
  box-sizing: border-box;
}
*,
*:before,
*:after {
  box-sizing: inherit;
}
```

[modern-normalize](https://github.com/sindresorhus/modern-normalize) already applies this.

### [Reset margin and padding](https://css-tricks.com/reset-all-margins-padding/)

```css
* {
  margin: 0;
  padding: 0;
}
```

### [text-rendering](https://css-tricks.com/almanac/properties/t/text-rendering/)

```css
p.legibility {
  text-rendering: optimizeLegibility;
}
p.speed {
  text-rendering: optimizeSpeed;
}
/* maybe not performant */
body {
  text-rendering: optimizeLegibility;
}
```

### [The 1140 grid fits perfectly into a 1280 monitor](https://www.ramotion.com/agency/web-design/cssgrid/)

```css
.row {
  max-width: 1140px;
}
```

### Align icons and text

```css
.icon-small {
  display: inline-block;
  width: 30px;
  text-align: center;
  font-size: 120%;
  margin-right: 10px;
  /* trick */
  line-height: 120%;
  vertical-align: middle;
  margin-top: -5px;
}
```

### Fixed background image (cool effect)

```css
.section-with-bg-image {
  background-image: linear-gradient(rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.7)),
    url(img/bg.jpg);
  background-size: cover;
  color: white;
  background-attachment: fixed;
}
```

### White box with hollow text over a background image

```css
h1.white {
  background: white;
  color: black;
  mix-blend-mode: screen;
}
```

### Black box with hollow text over a background image

```css
h1.black {
  background: black;
  color: white;
  mix-blend-mode: multiply;
}
```

### Title with image as text color

```css
h1.image {
  background-image: url(//unsplash.it/800);
  background-size: cover;
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}
```

### Center text on page

```css
body {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
```

### Render QR or thumbnails

```css
img {
  image-rendering: pixelated;
}
```

### Fluid font size

```css
:root {
  font-size: calc(1vw + 1vh + 0.5vmin);
}
```

### Clearfix

```css
/* old way */
.clearfix {
  overflow: auto;
}
/* new way */
.clearfix::after {
  content: '';
  clear: both;
  display: table;
}
/* future way */
.clearfix {
  display: flow-root;
}
```

### H1 with logo (SEO)

```html
<h1 class="logo">My Company and what we do</h1>
```

```css
.logo {
  position: relative;
  height: 50px;
  width: 50px;
  overflow: hidden;
}
.logo:after {
  content: url(logo.jpg);
  position: absolute;
  top: 0;
  left: 0;
}
```

### Image tips

- img width and height in percentage units does not work if img is inside an inline element. Wrapping element has to be changed to inline-block
- if img does not fit well in the wrapping div (white stripe at the bottom), use vertical-align: bottom or modify the wrapping tag to be displayed as inline-block or block.

### Responsive design without queries (using grid)

```css
main {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(19rem, 1fr));
  /* grid-gap: 1rem; */
}
```

### Block element only properties

- width
- margin: auto

### Screen readers only content

```css
.sr-only {
  border: 0;
  clip: rect(0 0 0 0);
  clip-path: polygon(0px 0px, 0px 0px, 0px 0px);
  -webkit-clip-path: polygon(0px 0px, 0px 0px, 0px 0px);
  height: 1px;
  margin: -1px;
  overflow: hidden;
  padding: 0;
  position: absolute;
  width: 1px;
  white-space: nowrap;
}
```

### Shorter way to make responsive width

```css
width: 70%;
max-width: 500px;
```

can be done shorter way:

```css
width: min(500px, 70%);
```

Besides min(), there are also max() and clamp() functions.

### Problem with no background color after inside elements set to float

This happens because a wrapping block element collapses when all inside block elements are float (to left or right). Solution for this is to set `overflow: auto` to the parent container.

### Properly display image in any rectangle area (show image central area not disturbing aspect ratio)

```css
object-fit: cover;
```

### Push footer to the page bottom

On a short page, html actually may not occupy the whole viewport height so when using percentages, it is important to use them for html element as well. Using percentages in this case gives better compatibility with smaller devices compared to wh.

```css
html {
  height: 100%;
}
body {
  min-height: 100%;
}
footer {
  margin-top: auto;
}
```

### Common bootstrap

```css
@import url(//cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css);

:root {
  font-size: calc(1vw + 1vh + 0.5vmin);
}

html {
  box-sizing: border-box;
}

*,
*:before,
*:after {
  box-sizing: inherit;
}

* {
  margin: 0;
  padding: 0;
}

body {
  height: 100vh;
  /* max-width: 1140px; */
  display: flex;
  justify-content: center;
  align-items: center;
  text-rendering: optimizeLegibility;
}
```

### Selective load of CSS (HTML <link> media Attribute)

```css
<link rel="stylesheet" type="text/css" href="print.css" media="print">
<link rel="stylesheet" type="text/css" href="mobile.css" media="screen and (max-width: 768)">
```

### Smooth Scroll-to-top Button

```css
html {
  scroll-behavior: smooth;

  @media (prefers-reduced-motion: reduce) {
    scroll-behavior: auto;
  }
}
```

```html
<button
  onClick={() => {
    window.scrollTo(0, 0);

    // focus management
    const title = document.getElementById('article-title');
    title.focus();
  }}
>
  Scroll to top
</button>
```
