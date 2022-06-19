# Frontend - JavaScript, HTML5, CSS3, resources

- [CSS](css.md)
- [JavaScript & TypeScript](javascript-typescript.md)

## Documentation

- [Frontend Book](https://frontendmasters.com/books/front-end-handbook/2019/#1)
- [ES6 Cheatsheet](https://devhints.io/es6)
- [ES6 Arrow Functions](https://www.sitepoint.com/es6-arrow-functions-new-fat-concise-syntax-javascript/)
- [Drag & Drop With Vanilla JS](https://www.youtube.com/watch?v=C22hQKE_32c&t=17s)
- [Dark theme in a day](https://medium.com/@mwichary/dark-theme-in-a-day-3518dde2955a)
- [prefers-color-scheme media query](https://web.dev/prefers-color-scheme/)
- [Dark Mode in CSS - prefers-color-scheme](https://css-tricks.com/dark-modes-with-css/)
- [YouTube Player API Reference for iframe Embeds](https://developers.google.com/youtube/iframe_api_reference)
- [window.location cheatsheet](https://dev.to/samanthaming/window-location-cheatsheet-4edl)
- [40 Free HTML landing page templates](https://dev.to/davidepacilio/40-free-html-landing-page-templates-3gfp)
- [Easing funtions](https://easings.net/en)
- [Browser Default Styles](https://browserdefaultstyles.com/)
- [Spread Syntax Three-dots - Tricks You Can Use Now](https://dev.to/girlie_mac/spread-syntax-three-dots-tricks-you-can-use-now-aob)
- [When to use SVG and when to use Canvas](https://css-tricks.com/when-to-use-svg-vs-when-to-use-canvas/)

## Design

- [Colorinspo - all in one resource to find everything about colors with extreme ease](https://colorsinspo.com/)
- [Color Hunt - Dark Color Palettes](https://colorhunt.co/palettes/dark)
- [CSS Gradient](https://cssgradient.io/)
- [Make a smooth shadow](https://brumm.af/shadows)
- [Flat UI Colors](https://flatuicolors.com/)
- [Material Flat UI Colors](https://www.materialui.co/flatuicolors)
- [Social Media Colors 2020](https://www.lockedownseo.com/social-media-colors/)
- [ColorSpace](https://mycolor.space/)
- [0to255](https://www.0to255.com/)
- [colors.lol](https://colors.lol/)
- [ColorMind - color scheme generator](http://colormind.io/)
- [Gradient Magic - Free Gallery of Fantastic and Unique CSS Gradients](https://www.gradientmagic.com/)
- [SVG Backgrounds](https://www.svgbackgrounds.com/)
- [freepik Background Collection](https://www.freepik.com/free-photos-vectors/svg-background)
- [Hero Patterns - collection of repeatable SVG background patterns](https://www.heropatterns.com/)
- [flaticon - design background of icons](https://www.flaticon.com/pattern/)
- [Loading Patterns](https://loading.io/pattern/)
- [khroma - design with colors you love](http://khroma.co/)
- [coolors - super fast color schemes generator](https://coolors.co/)
- [Get Waves - create SVG waves](https://getwaves.io/)
- [Gradient Hunt - platform for color inspiration with thousands of trendy hand-made color gradients](https://gradienthunt.com/)

## Design Showcases

- [dribbble](https://dribbble.com/)
- [Behance](https://www.behance.net/)
- [siteinspire](https://www.siteinspire.com/)
- [awwwards](https://www.awwwards.com/)
- [Land-book](https://land-book.com/)

## Illustrations

- [error404](https://error404.fun/)
- [Blush](https://blush.design/)
- [Smash Illustrations](https://usesmash.com/)
- [Control](https://control.rocks/)
- [DrawKit](https://www.drawkit.io/)
- [Open Doodles](https://www.opendoodles.com/)
- [Free Illustrations](https://freellustrations.com/)
- [Mixkit](https://mixkit.co/free-stock-art/)
- [Delesign](https://delesign.com/free-designs/graphics/)
- [unDraw](https://undraw.co/illustrations)

## Stock Images

- [Pexels](https://www.pexels.com)
- [Unsplash](https://unsplash.com)
- [picjumbo](https://picjumbo.com)
- [pixabay](https://pixabay.com)
- [AllTheFreeStock - links to multiple resources](https://allthefreestock.com)
- [Lorem Picsum - Example photos - unsplash.it](https://picsum.photos/)

## Icons

- [Material Design Icons By Community](https://materialdesignicons.com/)
- [Font Awesome](https://fontawesome.com/)
- [iconogram - all icons as separate .svg files](https://icongr.am/)
- [IcoMoon](https://icomoon.io/app/#/select)
- [Genericons](https://github.com/Automattic/genericons-neue)
- [Social Logos](https://github.com/Automattic/social-logos)
- [Linea Icon](https://linea.io)
- [Free SVG icons for popular brands](https://simpleicons.org)
- [SVGBox - copy and paste icons](https://svgbox.net/iconsets/)
- [SVG Silh - free SVG images & icons](https://svgsilh.com/)
- [Streamline Icons](https://app.streamlinehq.com/icons)
- [Tabler Icons]([Algorithms for Decision Making](https://tabler-icons.io/))
- [heroicons - beautiful hand-crafted SVG icons by the makers of Tailwind CSS](https://heroicons.com/)

## Fonts

- [Google Fonts](https://fonts.google.com/)
- [Overpass](http://overpassfont.org/)
- [Overpass Cyrillic](https://github.com/AlexxNB/overpass-font-cyrillic)
- [glyphhanger - subsets web fonts](https://github.com/zachleat/glyphhanger)
- [Font style matcher](https://meowni.ca/font-style-matcher/)
- [FontSpark allows you to discover your next favorite font](https://fontspark.app/)
- [Web Safe Fonts](https://www.cssfontstack.com/)
- [Nunito](https://fonts.google.com/specimen/Nunito)
- [GDPR Compliant Replacement for Google Fonts](https://fonts.bunny.net/about)

### Replace Google fonts with GDPR compliant mirror

https://fonts.googleapis.com/css -> https://fonts.bunny.net/css

## Security

- [Protect your email address from spam](https://www.email-encoder.com/)
- [Understanding CSRF](https://github.com/pillarjs/understanding-csrf)

### Generate SHA384

```bash
wget https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css

cat font-awesome.min.css | openssl dgst -sha384 -binary | openssl enc -base64
```

Link:

```html
<link
  rel="stylesheet"
  href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
  integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN”
  crossorigin=“anonymous"
/>
```

## Editors

- [Editor.js](https://editorjs.io/)
- [Rendering](https://github.com/codex-team/codex/tree/master/www/application/views/templates/editor/plugins)
- [CKEDitor 5](https://ckeditor.com/ckeditor-5/demo/)
- [tiptap](https://github.com/scrumpy/tiptap)
- [Quill - powerful rich text editor](https://quilljs.com/)
- [Slate - completely customizable framework for building rich text editors](https://github.com/ianstormtaylor/slate)

## Upload

- [Amazon S3 Browser Upload](https://www.shanestillwell.com/2018/09/02/amazon-file-upload/)
- [Secure File Upload to S3 Directly from the Browser](https://m.youtube.com/watch?v=shCCP4PFXeU)

## Mobile Apps

- [NativeScript - Create Native iOS and Android Apps with JavaScript](https://www.nativescript.org)
- [Weex](https://weex.apache.org/)

## HTML5

- [20 Days of HTML](https://dev.to/theindiancodinggrl/20-days-of-html-learn-20-amazing-things-about-html-part-1-2p99)
- [<summary> and <details> - Creates a disclosure widget in which information is visible only when the widget is toggled into an "open" state](https://www.w3schools.com/tags/tag_summary.asp)

### Components

- [Toggle Switch](https://codepen.io/mburnette/pen/LxNxNg)

## SEO

- [Open Graph protocol enables any web page to become a rich object in a social graph (Facebook)](https://ogp.me/)
- [Facebook and Twitter Cards Validator](https://www.opengraph.xyz/)
- [Twitter Card](https://developer.twitter.com/en/docs/tweets/optimize-with-cards/guides/getting-started)
- [Twitter Card Validator](https://cards-dev.twitter.com/validator)
- [iframely - Meta Validator](http://debug.iframely.com/)

## Accessibility

- [A11Y Project](https://a11yproject.com/patterns/)
- [WAI-ARIA Recommendation](https://www.w3.org/TR/wai-aria-1.1/)

## Fake backends

- [JSONPlaceholder](https://jsonplaceholder.typicode.com)
- [DOG CEO](https://dog.ceo/)

## Maps

- [Free SVG countries' vector maps](https://mapsvg.com/maps)

## Camera, video and screen streaming

- [Stream Webcam to Java Server with Websocket](https://www.youtube.com/watch?v=H42bl4RDQF8)
- [MediaDevices](https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices)
- [Screen Capture API](https://developer.mozilla.org/en-US/docs/Web/API/Screen_Capture_API)
- [Introduction to getUserMedia](https://m.youtube.com/watch?v=Hc7GE3ENz7k)

## Tools

- [SVGOMG - GUI for optimizing SVG files](https://jakearchibald.github.io/svgomg/)
- [ShrinkMe](https://shrinkme.app/)
- [Squoosh - Compress your images will almost unnoticeable quality loss](https://squoosh.app/)
- [Lighthouse - analyzes web apps and web pages, collecting modern performance metrics and insights](https://github.com/GoogleChrome/lighthouse)
- [Responsive Image Breakpoints Generator](https://responsivebreakpoints.com/)
- [Favicon Generator - All browsers, all platforms](https://realfavicongenerator.net/)
- [Optimizilla - Online image optimizer](https://imagecompressor.com/)
- [W3C Markup Validation Service](https://validator.w3.org/)
- [W3C CSS Validation Service](https://jigsaw.w3.org/css-validator/)
- [Test mobile site](https://www.thinkwithgoogle.com/intl/en-gb/feature/testmysite/)
- [Majestic - Zero config GUI for Jest](https://github.com/Raathigesh/majestic)
- [carbon - Create and share beautiful images of your source code](https://carbon.now.sh/)
- [Wappalyzer - Find out stack of your favorite website](https://www.wappalyzer.com/)
- [Shape Divider - Generate sleek dividers for your website](https://www.shapedivider.app/)
- [Ucraft Logo Maker](https://www.ucraft.com/free-logo-maker)
- [AppMockUp - Generate mock-ups for Android and iPhones](https://app-mockup.com/)
- [Webframe - Thousands of design inspiration based on real websites](https://webframe.xyz/)
- [Remove BG - Delete the background from your images](https://www.remove.bg/)
- [Unscreen - remove.bg for gifs and videos](https://www.unscreen.com/)
- [Muffet - performant website link checker](https://github.com/raviqqe/muffet)

## Performance

- [GTmetrix](https://gtmetrix.com/)
- [Google PageSpeed Insights](https://developers.google.com/speed/pagespeed/insights/)
- [Lighthouse](https://developers.google.com/web/tools/lighthouse)
