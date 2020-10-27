# Angular

- https://angular.io
- https://cli.angular.io
- https://augury.angular.io

- AngularJS rewritten to Angular2
- New release every 6 months, now at Angular10
- Mostly backwards compatible

---

## CLI fundamentals

```shell script
# create a new app
$ ng new app-name

# create a component
$ ng generate component component-name
# or
$ ng g c component-name

# serve the app
$ ng serve
```

---

## Loading and start up

- `js` injected into index.html
- `main.ts` imports the main module - `app.module.ts`
- `AppModule` decorator received a list of components for bootstrapping, starting with `AppComponent`
- `AppComponent` injects itself into the specified selector in`index.html`, eg `<app-root></app-root>`

---

## AppModule and Component Declarations

- Most projects generally have just the one main module - `app.module.ts`
- New components must be _registered_ in the `@NgModule` decorator, `declarations` property
- `@NgModule` `imports` property lists other modules that are _imported_ into this main app module

---

## Components

- create with cli: `ng g c name`, or nested with `ng g c dir/name`
- [`@Component`](https://angular.io/api/core/Component) decorator sets up the components attributes
    - `selector` - css selector by: `element-name`, `[attribute-name]`, `.class-name`, _not_ by css id
    - `templateUrl` - path to html template file, can also use `template` for inline html template
    - `styleUrls` - an array of CSS files, can also use `styles` for inline css

---    
    
## Data binding

- Communication between component (typsecript code) and the template (html code)
- Output: component -> template
    - String interpolation: `{{ expression }}`
    - [Property binding](https://angular.io/guide/property-binding): `[property] = "expression"`
- Input: template -> component
    - [Event binding](https://angular.io/guide/event-binding): `(event) = "expression"`
- Bi-directional: component <-> template
    - [Two-way binding](https://angular.io/guide/two-way-binding) does both of the above, ie sets an element property and listens for an element change event
    - `[(ngModel)] = "data"` - remember as _banana in a box_, `[()]`    

To see a list of `properties` or `events` for a particular element can `console.log(element)`, also see 
MDN [HTML attribute reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes) and 
[Event reference](https://developer.mozilla.org/en-US/docs/Web/Events).

---

## Directives

- Directives are instructions in the DOM.
- Putting content into a location specified by a selector, eg `<app-foo></app-foo>` is a directive that includes a template
- Directives without templates are also possible, for example:

```typescript
@Directive({
    selector: '[appGreenText]'
})
export class GreenTextDirective{
    // ...
}
```

Which would be used like this:

```angular2html
<p appGreenText>This would be green</p>
``` 

### Built-in directives

- `ngIf` - used with `*` to indicate it is a _structural_ directive, ie `*ngIf`, eg
```angular2html
<p *ngIf="boolExpression">P tag shown if true</p>
```

- `ngStyle` - apply a css style, eg:
- Note the property binding is not part of the directive 
```angular2html
<p [ngStyle]="{'background-color': getColour()}"></p>
<p [ngStyle]="{backgroundColor: getColour()}"></p>
```

- `ngClass` - apply a css class, eg:
```angular2html
<p [ngClass]="{'class-name': boolExpression}"></p>
<p [ngClass]="{className: boolExpression}"></p>
```

- `ngFor` - structural directive for loops, eg
```angular2html
<div *ngFor="let item of itemList">{{ item.property }}</div>
<! -- with index -->
<div *ngFor="let item of itemList; let i = index">{{ i }} - {{ item.property }}</div>
```

---

## Debugging with Developer Console

- In developer tools, under **sources** tab
- When running in developer mode source mapping is used to link bundled code to the original source files
- Under `localhost:4200`, `main.js` - clicking any line number will open source file
- All source files are also located in the `webpack://` section
- Can add breakpoints etc
- Chrome extension _Augury_ can be added to dev tools 

---

## Binding to Custom Properties

- Allows data to be passed _into_ a component
- `@Input()` decorator is used to expose a class field such that it can be bound to properties from enclosing 
components, ie. passed in as 'props' in Vue parlance, eg:

```typescript
// Parent component
import {Component} from '@angular/core';
@Component({ selector: 'app-parent' })
export class ParentComponent {
  nameFromParent = 'Maia'; // Want this to be passed to child component
}

// Child component
@Component({ selector: 'app-child'})
export class ChildComponent {
  @Input() name: string; // decorator makes this a 'prop' 
}
```

In the `parent` component, where `nameFromParent` is available, the `name` field can be bound to `nameFromParent`:
```angular2html
<app-child [name]="nameFromParent"></app-child>
```

Then,  in the `child` component `name` will have the value of `nameFromParent`:

```angular2html
<div>{{ name }}</div>
```

- **Assigning an alias** for the bound property is also possible, by passing an arg to `@Input()`:

```angular2
// Child component
@Component({ selector: 'app-child'})
export class ChildComponent {
  @Input('childName') name: string; // now need to bind to childName in template 
}
```

`parent.component.html` binds to the aliases property:
```angular2html
<app-child [childName]="nameFromParent"></app-child>
```

`child.component.html` still uses the local property name:

```angular2html
<div>{{ name }}</div>
```

---

## Binding to Custom Events

- Allows data to be passed _out of_ (emitted from) a component
- `Output()` decorator marks a class field as an output property
- `Output('aliasName)` will assign an alias to the property

`child.component.ts`:
```angular2
@Component({ selector: 'app-child'})
export class ChildComponent {
  message: string;
  @Output() speak = new EventEmitter<string>();
  onSpeak() {
    this.speak.emit(this.message);
  }
}
```

`child.component.html`:
```angular2html
<div>
  <input type="text" [(ngModel)]="message">
  <button (click)="onSpeak()">Speak!</button>
</div>
```

---

## View Encapsulation

- CSS styles defined in the scope of a component are applied only to that component
- Angular ensures this by adding the same attribute to all elements in a component and applying styles to that attribute

Eg:

`some.component.css`
```css
p { color: red; }
```

`some.component.html`
```angular2html
<p>Style me!</p>
```

The source will look like: 
```html
<p _ngcontent-abc-123></p>
```

```css
p[_ngcontent-abc-123] { color: red; }
```

- [View encapsulation](https://angular.io/api/core/ViewEncapsulation) can be turned of at the component level, in the `@Component` decorator:

```angular2
import { ViewEncapsulation } from '@angular/core'
@Component({
    encapsulation: ViewEncapsulation.None
})
```

---

## Local References in Templates

- Alternative to to two-way binding
- Can be used on _any_ HTML element
- Syntax is `#arbitrayName`
- Creates a reference to the complete HTML element, not just its value
- The scope of this var is ONLY in the template, not in typsecript code
- Useful when the value just needs to be passed in from template, eg an input

```angular2html
<input type="text" #fooBar>
<button (click)="onClick(fooBar)">go</button>
```

```angular2
onClick(input: HTMLInputElement) {
    console.log(input.value)
}
```

- [`@VueChild()`](https://angular.io/api/core/ViewChild) decorator provides another way to access properties in the template
- It takes an argument which is the selector 
- Creates an `ElementRef` type
- Both `ViewChild` and `ElementRef` must be imported from `@angular/core`

```angular2html
<input type="text" #fooBar>
<button (click)="onClick()">go</button>
```

```angular2
import { ViewChild, ElementRef } from '@angular/core'
// ... //
class Foo{
    @ViewChild('fooBar') fooBar: ElementRef;
    onClick() {
        console.log(this.fooBar)
        console.log(this.fooBar.nativeElement.value)
    }
}
```

- **Note:** Should not change DOM elements using this method.

---

## Projecting Content into Components with `ng-content`

- `ng-content` is a directive (a hook) used to pass more complex HTML into a child component
- By default, anything added between opening and closing tags of an own component is ignored
- If `<ng-content></ng-content>` (nothing between) is located in a component template, the content between the opening 
and closing tags of that component, will be rendered.

`some.component.html`
```angular2html
<app-foo>
  <p>Hello!</p>
</app-foo>
``` 

`foo.component.html`
```angular2html
<p>The content between <app-foo></app-foo> will appear below:</p>
<ng-content></ng-content>
```

---

## [Component Lifecycle Hooks](https://angular.io/guide/lifecycle-hooks)

Note that some of these hooks can be triggered frequently so can affect performance. 

- `ngOnChanges()` - called on startup, and whenever a _bound_ property changes, properties with `@Input`
- `ngOnIt()` -  called once component is initialised, runs after constructor
- `ngDoCheck()` - called whenever change detection is run, which is on anything significant where a change is possible 
- `ngAfterContentInit()` - called after content (`ng-content`) has been projected into view
- `ngAfterContentChecked()` - called every time _projected_ content has been checked
- `ngAfterViewInit()` - called after the component's own view, and child views, have been initialised
- `ngAfterViewChecked()` - called when the component's own view, and child views, have been checked
- `ngOnDestroy()` - Called when component is about to be destroyed




 
  



