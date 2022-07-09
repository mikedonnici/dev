# Angular

- https://angular.io
- https://cli.angular.io
- https://augury.angular.io

- AngularJS rewritten to Angular2
- New release every 6 months, now at Angular10
- Mostly backwards compatible


- [CLI fundamentals](#cli-fundamentals)
- [Loading and start up](#loading-and-start-up)
- [AppModule and component declarations](#appmodule-and-component-declarations)
- [Components](#components)
- [Data binding](#data-binding)
- [Directives](#directives)
    - [Built-in directives](#built-in-directives)
    - [Custom directives](#custom-directives)
- [Debugging with developer console](#debugging-with-developer-console)
- [Binding to custom properties](#binding-to-custom-properties)
- [Binding to custom events](#binding-to-custom-events)
- [View encapsulation](#view-encapsulation)
- [Local references in templates](#local-references-in-templates)
- [Accessing the template and DOM with @ViewChild](#accessing-the-template-and-DOM-with-@viewchild)
- [Projecting content into components with `ng-content`](#projecting-content-into-components-with-ng-content)
- [Accessing `ng-content` with @ContentChild](#accessing-ng-content-with-@contentchild)
- [Component lifecycle hooks](#component-lifecycle-hooks)
- [Services and dependency injection](#services-and-dependency-injection)
    - [Hierarchical injection](#hierarchical-injection)
    - [Injecting services into services](#injecting-services-into-services)
- [Routing](#routing)
    - [Navigating programmatically](#navigating-programmatically)
    - [Route parameters](#passing-route-parameters)
    - [Query Parameters](#query-parameters)

---

## CLI fundamentals

```bash script
# create a new app
$ ng new app-name

# create a component
$ ng generate component component-name

# ...short version
$ ng g c component-name

# create a component without testing code
$ ng g c component-name --spec false

# serve the app
$ ng serve
```

---

## Loading and start up

- `js` injected into index.html
- `main.ts` imports the main module - `app.module.ts`
- `AppModule` decorator received a list of components for bootstrapping,
  starting with `AppComponent`
- `AppComponent` injects itself into the specified selector in`index.html`,
  eg `<app-root></app-root>`

---

## AppModule and component declarations

- Most projects generally have just the one main module - `app.module.ts`
- New components must be _registered_ in the `@NgModule`
  decorator, `declarations` property
- `@NgModule` `imports` property lists other modules that are _imported_ into
  this main app module

---

## Components

- create with cli: `ng g c name`, or nested with `ng g c dir/name`
- [`@Component`](https://angular.io/api/core/Component) decorator sets up the
  components attributes
    - `selector` - css selector by: `element-name`, `[attribute-name]`
      , `.class-name`, _not_ by css id
    - `templateUrl` - path to html template file, can also use `template` for
      inline html template
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
    - [Two-way binding](https://angular.io/guide/two-way-binding) does both of
      the above, ie sets an element property and listens for an element change
      event
    - `[(ngModel)] = "data"` - remember as _banana in a box_, `[()]`

To see a list of `properties` or `events` for a particular element
can `console.log(element)`, also see
MDN [HTML attribute reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes)
and
[Event reference](https://developer.mozilla.org/en-US/docs/Web/Events).

---

## Directives

- Directives are instructions in the DOM
- Two types:
    - **Attribute Directives**
        - look like an HTML attribute, can have data or event binding
        - only affect elements they are added to
    - **Structural Directives**
        - look like an HTML attribute but with leading `*` (for desugaring)
        - affect an area of the DOM as elements can be added or removed
        - Can only have one structural directive on an element
- Putting content into a location specified by a selector,
  eg `<app-foo></app-foo>` is a directive that includes a template

### Built-in directives

#### `ngIf` structural directive

```angular2html
<p *ngIf="boolExpression">P tag shown if true</p>
```

#### `ngFor` structural directive

```angular2html

<div *ngFor="let item of itemList">{{ item.property }}</div>
<! -- with index -->
<div *ngFor="let item of itemList; let i = index">{{ i }}
    - {{ item.property }}</div>
```

#### `ngStyle` attribute directive

- Note that the square brackets indicate a binding to some property on the
  `ngClass` directive and are not part of the directive itself.

```angular2html
<p [ngStyle]="{'background-color': getColour()}"></p>
<p [ngStyle]="{backgroundColor: getColour()}"></p>
```

#### `ngClass` attribute directive

```angular2html
<p [ngClass]="{'class-name': boolExpression}"></p>
<p [ngClass]="{className: boolExpression}"></p>
```

### Custom directives

- Typically create a folder and file for a custom directive, eg:
  `src/green-text/green-text.directive.ts`
- Use `@Directive` decorator and specify a selector:
    - Note that specifying the selector with square brackets indicates that the
      selector can be used as a stand-alone attribute on the element, ie in the
      form: `<p someSelector>Hello!</p>`
- Angular will instantiate the class and pass the ElementRef to `constructor()`
- Directive needs to be imported in `app.module.ts`

```bash
# Generate a directive
$ ng generate directive greenText 
# or ng g d greenText
```

```typescript
import {Directive, ElementRef, OnInit} from "@angular/core"

@Directive({
    selector: '[appGreenText]'
})
export class GreenTextDirective implements OnInit {

    constructor(private elementRef: ElementRef) {
    }

    ngOnInit() {
        this.elementRef.nativeElement.style.backgroundColor = 'green';
    }
}
```

- use like this:

```angular2html
<p appGreenText>This would be green</p>
```

The above works, however, accessing elements directly like this is _not_ good
practice.

A better way is to use the Renderer:

```typescript
import {Directive, ElementRef, OnInit, Renderer2} from '@angular/core';

@Directive({
    selector: '[appGreenText]'
})
export class GreenTextDirective implements OnInit {

    constructor(private elementRef: ElementRef, private renderer: Renderer2) {
    }

    ngOnInit(): void {
        this.renderer.setStyle(this.elementRef.nativeElement, 'background-color', 'green');
    }
}
```

`@HostListener()` decorator can be used to create a _reactive_ directive.

For example, to change text colour only on hover:

```typescript
@Directive({
    selector: '[appGreenText]'
})
export class GreenTextDirective {

    constructor(private elementRef: ElementRef, private renderer: Renderer2) {
    }

    @HostListener('mouseenter') hoverOn(_: Event) {
        this.renderer.setStyle(this.elementRef.nativeElement, 'background-color', 'green');
    }

    @HostListener('mouseleave') hoverOff(_: Event) {
        this.renderer.setStyle(this.elementRef.nativeElement, 'background-color', 'transparent');
    }
}
```

`@HostBinding()` decorator can also be used to achieve the above, in an even
simpler way:

```typescript
@Directive({
    selector: '[appGreenText]'
})
export class GreenTextDirective {

    @HostBinding('style.backgroundColor') backgroundColor: string = 'transparent';

    @HostListener('mouseenter') hoverOn(_: Event) {
        this.backgroundColor = 'green';
    }

    @HostListener('mouseleave') hoverOff(_: Event) {
        this.backgroundColor = 'transparent';
    }
}
```

---

## Debugging with developer console

- In developer tools, under **sources** tab
- When running in developer mode source mapping is used to link bundled code to
  the original source files
- Under `localhost:4200`, `main.js` - clicking any line number will open source
  file
- All source files are also located in the `webpack://` section
- Can add breakpoints etc
- Chrome extension _Augury_ can be added to dev tools

---

## Binding to custom properties

- Allows data to be passed _into_ a component
- `@Input()` decorator is used to expose a class field such that it can be bound
  to properties from enclosing components, ie. passed in as 'props' in Vue
  parlance, eg:

```typescript
// Parent component
import {Component} from '@angular/core';

@Component({selector: 'app-parent'})
export class ParentComponent {
    nameFromParent = 'Maia'; // Want this to be passed to child component
}

// Child component
@Component({selector: 'app-child'})
export class ChildComponent {
    @Input() name: string; // decorator makes this a 'prop' 
}
```

In the `parent` component, where `nameFromParent` is available, the `name` field
can be bound to `nameFromParent`:

```angular2html

<app-child [name]="nameFromParent"></app-child>
```

Then, in the `child` component `name` will have the value of `nameFromParent`:

```angular2html

<div>{{ name }}</div>
```

- **Assigning an alias** for the bound property is also possible, by passing an
  arg to `@Input()`:

```typescript
// Child component
@Component({selector: 'app-child'})
export class ChildComponent {
    // now need to bind to childName in template  
    @Input('childName') name: string;
}
```

`parent.component.html` binds to the aliases property:

```angular2html

<app-child [childName]="nameFromParent"></app-child>
```

Note that `child.component.html` still uses the local property name:

```angular2html

<div>{{ name }}</div>
```

---

## Binding to custom eventss

- Allows data to be passed _out of_ (emitted from) a component
- `Output()` decorator marks a class field as an output property
- `Output('aliasName)` will assign an alias to the property

`child.component.ts`:

```typescript
@Component({selector: 'app-child'})
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

## View encapsulation

- CSS styles defined in the scope of a component are applied only to that
  component
- Angular ensures this by adding the same attribute to all elements in a
  component and applying styles to that attribute

Eg:

`some.component.css`

```css
p {
    color: red;
}
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
p[_ngcontent-abc-123] {
    color: red;
}
```

- [View encapsulation](https://angular.io/api/core/ViewEncapsulation) can be
  turned of at the component level, in the `@Component` decorator:

```angular2
import { ViewEncapsulation } from '@angular/core'
@Component({
    encapsulation: ViewEncapsulation.None
})
```

---

## Local references in templates

- Alternative to to two-way binding
- Can be used on _any_ HTML element
- Syntax is `#arbitrayName`
- Creates a reference to the _complete HTML element_, not just its value
- The scope of this var is ONLY in the template, not in TS code
- Useful when the value just needs to be passed in from template, eg an input

```angular2html
<input type="text" #fooBar>
<button (click)="onClick(fooBar)">go</button>
```

```typescript
onClick(input
:
HTMLInputElement
)
{
    console.log(input.value)
}
```

## Accessing the template and DOM with @ViewChild

- [`@VueChild()`](https://angular.io/api/core/ViewChild) decorator provides
  another way to access properties in the template
- It takes an argument which is the selector
- Creates an `ElementRef` type
- Both `ViewChild` and `ElementRef` must be imported from `@angular/core`

```angular2html
<input type="text" #fooBar>
<button (click)="onClick()">go</button>
```

```typescript
import {ViewChild, ElementRef} from '@angular/core'

// ... //
class Foo {
    @ViewChild('fooBar', {static: true}) fooBar: ElementRef;

    onClick() {
        console.log(this.fooBar)
        console.log(this.fooBar.nativeElement.value)
    }
}
```

- **Note:** Should not change DOM elements using this method.

---

## Projecting content into components with `ng-content`

- `ng-content` is a directive (a hook) used to pass more complex HTML into a
  child component
- By default, anything added between opening and closing tags of a custom
  component is ignored
- If `<ng-content></ng-content>` (nothing between) is located in a component
  template, the html between the opening and closing tags where that component
  is used, will be rendered.

`some.component.html`

```angular2html

<app-foo>
    <p>Hello!</p>
</app-foo>
``` 

`foo.component.html`

```angular2html

<p>The content between
    <app-foo></app-foo>
    will appear below:
</p>
<ng-content></ng-content>
```

---

## Accessing `ng-content` with @ContentChild

- Provides a way to access local template references in the _parent_ component

For example:

```angular2html
<p>This is app-parent</p>
<app-child>
    <p>
        This content will appear in app-child at the position of ng-content
    </p>
    <p #aParagraph>
        This paragraph is referenced by a local template var: aParagraph. To
        access
        this in app-child we could use @ViewChild. BUT, to access this paragraph
        in app-parent we can use @ContentChild.
    </p>
</app-child>
```

```angular2html
<p>This is app-child</p>
<ng-content></ng-content>
```

```typescript
class ParentComponent {
    // This provides access to a local reference that appears in the template 
    // between the opening and closing tags of a component.
    // Note that this is not available until after ngOnInit().
    @ContentChild('aParagraph') aParagraph: ElementRef;

    ngOnInit(): void {
        console.log(this.aParagraph.nativeElement.textContent) // -> empty 
    }

    ngAfterInit(): void {
        console.log(this.aParagraph.nativeElement.textContent) // -> OK
    }
}
```

---

## [Component lifecycle hooks](https://angular.io/guide/lifecycle-hooks)

Note that some of these hooks can be triggered frequently so can affect
performance.

- `ngOnChanges()` - called on startup, and whenever a _bound_ property changes,
  properties with `@Input`
- `ngOnIt()` - called once component is initialised, runs after constructor
- `ngDoCheck()` - called whenever change detection is run, which is on anything
  significant where a change is possible
- `ngAfterContentInit()` - called after content (`ng-content`) has been
  projected into view
- `ngAfterContentChecked()` - called every time _projected_ content has been
  checked
- `ngAfterViewInit()` - called after the component's own view, and child views,
  have been initialised
- `ngAfterViewChecked()` - called when the component's own view, and child
  views, have been checked
- `ngOnDestroy()` - Called when component is about to be destroyed

## Services and dependency injection

- Classes used to centralise functionality and communicate between components
- Can generate with cli: `ng generate service foo` or `ng g s foo`
- For example, a simple logging service `logging.service.ts`:

```typescript
export class LoggingService {
    logStatusChange(status: string) {
        console.log("Status changed to " + status)
    }
}
```

- Dependency injection is used to access services from a component
- Objects of the required service classes will be automatically _injected_ into
  the components constructor as determined by the argument types
- For example, to access the `LoggingService` from a component:
    - Add `providers` to `Component` decorator arg
    - Add service param to the component constructor

```typescript
import {Component} from '@angular/core'
import {LoggingService} from 'logging.service'

@Component({
    // ... //
    providers: [LoggingService]
})
export class FooComponent {

    // Service param added here...
    constructor(logger: LoggingService) {
    }

    // service object available as a property
    someFunc() {
        this.logger.logStatusChange('bla bla')
    }
}
```

### Hierarchical injection

- The same instance of a service object is available from the point of injection
  and at every point below, in the component hierarchy:

| Injection Site  | Availability                                         |
|-----------------|------------------------------------------------------|
| `AppModule`     | **application-wide** , ie all components and services|
|`AppComponent`   | **all components**, but nout other services          |
| other component | Same component and all child components              |

- Important to note that instantiating a service object will override any
  instances that were created at a higher level.
- The `providers` array in the `@Component` decorator specifies which service
  objects will be _instantiated_.
- Hence, if a service object already exists, and the same object is required,
  that service should be _omitted_ from the `providers` array, but maintained as
  an argument to the component constructor:

```typescript
// In this component we need to access Service1 and Service2
// Service1 was instantiated somewhere above this point in the hierarchy
@Component({
    providers: [Service2] // Note: Service1 omitted
})
export class FooComponent {
    // Both services added to constructor, Service2 will be instantiated here
    constructor(private srvc1: Service1, private srvc2: Service2) {
    }
}
```

### Injecting services into services

- If a service is added to the `providers` array in `AppModule` it is available
  _application-wide_
- The `@Injectible()` decorator must then be added to any service that is to
  _receive_ the instance of this top-level service.
- Note that even though the `Injectible()` decorator is only required on
  service _receiving_ another service, it is now recommended practice to add
  the `@Injectible()` decorator to service that are _being_ injected.

```typescript
@NgModule({
    //...//
    providers: [Service1], // now available everywhere
    bootstrap: [AppComponent]
})
export class AppModule {
}
```

```typescript
@Injectible() // can now have other services injected
export class Service2 {
    constructor(srvc1: Service1) {
    }
}
```

- From Angular 6 onwards there is another way to achieve the same as above.
- Provide an arg to `@Injectible()` and the service does not have to be added
  to `providers` in `AppModule`:

```typescript
@Injectible({providedIn: 'root'})
export class Service2 {
    constructor(srvc1: Service1) {
    }
}
```

- This allows Angular to _lazy load_ code and may result in better performance
  for larger apps.
  
---

## Routing

- Routes are set up and registered in `app.module.ts`
- `RouterModule.forRoot()` registers the routes for the main application

```typescript
// Set up routes
import {RouterModule, Routes} from '@angular/router';

const appRoutes: Routes = [
    {path: '', component: HomeComponent},
    {path: 'dog', component: DogComponent},
    {path: 'cat', component: CatComponent},
];

// register routes
@NgModule({
    // ... //
    imports: [
        RouterModule.forRoot(appRoutes),
    ]
})
```

- The location for the output for the currently selected route is specified with
  the `router-outlet` directive:

```angular2html

<router-outlet></router-outlet>
```

- Linking to routes with normal `href` creates a new request which reloads the
  page and loses state.
- Use `routerLink` to specify path or bind to `[routerLink]` to create more
  complex paths (`["a", "b", "c"]` = `/a/b/c`):

```angular2html
<a routerLink="/">Home</a>
<a routerLink="/parent">Parent</a>
<a [routerLink]="['/parent', 'child']">Child</a>
```

- Note that route paths without a leading `/` will be relative to _current path_
- Can also use dir path notation, eg `./same-level`, `../up-one`

- `routerLinkActive` can be used to assign a css class to the currently active
  route:

```angular2html

<li routerLinkActive="some-class"><a routerLink="/">Home</a></li>
<li routerLinkActive="some-class"><a routerLink="/foo">Foo</a></li>
```

- However, `routerLinkActive` matches the path from left to right, so in the
  example above `/` would match both home and foo.
- To alleviate this the `routerLinkActiveOptions` is used.
- This will mark the route as active based on the _entire_ path, as opposed to
  the first match from the left

```angular2html

<li routerLinkActive="some-class" [routerLinkActiveOptions]="{exact: true}">
    <a routerLink="/">Home</a>
</li>
```

### Navigating programmatically

- Inject the router into a component constructor:

```typescript
import {Router} from "@angular/core"

export class SomeComponent {

    constructor(private router: Router) {
    }

    onSomething() {
        this.router.navigate(["/to", "this", "path"])
    }
}
```

- for relative paths need access to the current route, via `ActivatedRoute`, and
  pass a second arg to `.navigate()`:

```typescript
import {Router, ActivatedRoute} from "@angular/core"

export class SomeComponent {

    constructor(private router: Router, private route: ActivatedRoute) {
    }

    onSomething() {
        this.router.navigate(["child"], {relativeTo: this.route})
    }
}
```

### Route parameters

- Use a colon in route path to designate a path variable, eg `/users/:id`
- To access the variable inject `ActivatedRoute` into the component constructor
  and access via `.snapshot.params[name], eg:

```typescript
import {ActivatedRoute} from "@angular/core"

export class UserComponent {
    user: { id: number, name: string }

    constructor(private route: ActivatedRoute) {
    }

    ngOnInit() {
        this.user = {
            id: this.route.snapshot.params['id'],
            name: getUserName(this.route.snapshot.params['id'])
        }
    }
}
```

- The approach above will only modify the `user` object when the compnent is
  reloaded
- If the url params are changed from within the _current_ component then a
  different approach is required
- To update url params _reactively_, need to subscribe to the `params`
  observable:

```typescript
import {ActivatedRoute} from "@angular/core"

export class UserComponent {
    user: { id: number, name: string }

    constructor(private route: ActivatedRoute) {
    }

    ngOnInit() {
        this.user = {
            id: this.route.snapshot.params['id'],
            name: getUserName(this.route.snapshot.params['id'])
        }
        this.route.params.subscribe((params: Params) => {
            this.user.id = params['id']
            this.user.name = getUserName(params['id'])
        })
    }
}
```

- In this case, Angular will take care of running`.unsubscribe()` to clean up
  memory when a component is destroyed. In some cases, such as custom
  observables you might need to run `.unsubscribe()` in the `onDestroy()`
  lifecycle hook.

### Query Parameters and Fragments

- Query params can be added to a router link by binding to `queryParams` 
  and passing it an object or key-value pairs.
- A fragment is added by binding to `fragment`.

```angular2html
<a
        [routerLink]="['/some', 'path']"
        [queryParams]="{key1: 'value1', key2: 10}"
        fragment="here"
>Foo</a>
```
- This is: `/some/path?key1=value1&key2=10#here`
- To do this programmatically:

```typescript
this.router.navigate(
    ['/some', 'path'], 
    {queryParams: {key1: 'value1', key2: 10}, fragment: "here"}
)
```

- Query params and fragment are also available on the `ActivatedRoute` which can 
  be injected into a module
- Values can be retrieved in a similar way to params, ie via the snapshot 
  or using observables:

```typescript
export class SomeComponent {
  ngOnInit() {
    console.log(this.route.snapshot.queryParams)
    console.log(this.route.snapshot.fragment)
    this.route.queryParams.subscribe()
    this.route.fragment.subscribe()
  }    
}
```

[incomplete]

- Navigating to routes should be done with the `routerlink` directive, rather
  than `href`, as the latter will reload the page and lose state.
- Each of the following will work:

```angular2html

```


```typescript

```
---

## Observables

- A pattern used to handle asynchronous tasks
- The **Observable** represents a source of data such as events, http requests
- The **Observer** subscribes to the _Observable_ and executes code that:
    - Handles data
    - Handles errors
    - Handles completion (where applicable)
- An alternative approach to promises
- Observable is an object imported from a 3rd=party package such as `rxjs`, that 
  follows the _observable_ pattern:
   - observable -> trigger emits data -> observer
- The _observer_ is you own code that _subscribes_ to the observable
- `rxjs` package provides multiple ways for creating an observable, eg a simple
  one is `interval`
  
```ts
import { Component, OnInit } from '@angular/core';
import {interval} from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'observables';

  ngOnInit(): void {
    interval(1000).subscribe((n) => {
      console.log(n);
    });
  }
}
```

- This will log an incrementing number to the console, indefinitely, so it is 
  important to `unsubscribe()` from observables in order to prevent memory leaks.
- Revisiting this component will also create a separate observable each time so 
  can also compound memory use.
- So would do something like:

```ts
import {Component, OnDestroy, OnInit} from '@angular/core';
import {interval, Subscription} from 'rxjs';

@Component({
  selector: 'app-obs',
  templateUrl: './obs.component.html',
  styleUrls: ['./obs.component.css']
})
export class ObsComponent implements OnInit, OnDestroy {

  private counterSubs = new Subscription();

  constructor() { }

  ngOnInit(): void {
    this.counterSubs = interval(3000).subscribe(n => {
      console.log(`Obs count = ${n}`);
    });
  }

  ngOnDestroy(): void {
    console.log('Unsubscribe()');
    this.counterSubs.unsubscribe();
  }
}
```

- Observables that are part of angular packages normally take care of 
  _unsubscribing_ for you, so don't generally have to worry about it.
  
- Building a _custom_ observable:

```ts
const customIntervalObservable = new Observable(observer => {
  let count = 0;
  setInterval(() => {
    count++;
    if (Math.floor(Math.random() * 5) === 3) {
      observer.error(new Error('1 in 5 chance of an error'));
    }
    if (count === 10) {
      observer.complete();
    }
    observer.next(count);
  }, 1000);
});

```

- When the observable throws an error it also stops emitting data and there 
  is no need to unsubscribe.
- Second argument to `subscribe()` is an error handling function
- Third arg to `subscribe()` is a function to handle the `.complete()` call

```ts
this.counterSubs = customIntervalObservable.subscribe(
      n => {
        console.log(`Obs2 count = ${n}`);
      },
      e => {
        alert(`Error: ${e.message}`);
      },
      () => {
        console.log('Complete');
      }
    );
```
- Rarely build your own observables and will usually use `subscribe()` by 
  passing in functions that deal with data, errors or completion.

### Operators

- Operators are a useful feature of the `rxjs` library that enable 
  transformation of the data being emitted by the observable.
- `rxjs` has many [built-in operators](https://rxjs-dev.firebaseapp.com/guide/operators)
- Observables have `pipe()` method which receive an operator function arg, 
  which itself receives the same _data_ that is passed to `subscribe()`
- Returns a new observable so need to subscribe to that to get the transformed 
  data:

```ts
// Using the map operator from rxjs
import {map} from 'rxjs/operators';
// ... //
const newObservable = originalObservable.pipe(map((data: any) => {
  return `Transformed ${data}`;
}));
newObservable.subscribe( n => {
  console.log(n);
});
```

- Note that `pipe()` can take multiple arguments which are sequentially applied
  to the data

### Subject

- A `Subject` is an `rxjs` class that can be used in place of an `EventEmitter`
- It is a special type of `Observable` that allows for `.next()` to be called 
  from _outside_ the observable, can be triggered from code.
- Should also `unsubscribe` from `Subjects` in `onDestroy`
- This mechanism is only really used across components via a service - for 
  events within a component would still use Angular's `EventEmitter`. 

  

  

