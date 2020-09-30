# Angular

- https://angular.io
- https://cli.angular.io
- https://augury.angular.io

- AngularJS rewritten to Angular2
- New release every 6 months, now at Angular10
- Mostly backwards compatible

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


## Loading and start up

- `js` injected into index.html
- `main.ts` imports the main module - `app.module.ts`
- `AppModule` decorator received a list of components for bootstrapping, starting with `AppComponent`
- `AppComponent` injects itself into the specified selector in`index.html`, eg `<app-root></app-root>`


## AppModule and Component Declarations

- Most projects generally have just the one main module - `app.module.ts`
- New components must be _registered_ in the `@NgModule` decorator, `declarations` property
- `@NgModule` `imports` property lists other modules that are _imported_ into this main app module

## Component

- create with cli: `ng g c name`, or nested with `ng g c dir/name`
- [`@Component`](https://angular.io/api/core/Component) decorator sets up the components attributes
    - `selector` - css selector by: `element-name`, `[attribute-name]`, `.class-name`, _not_ by css id
    - `templateUrl` - path to html template file, can also use `template` for inline html template
    - `styleUrls` - an array of CSS files, can also use `styles` for inline css
    
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


## Debugging with Developer Console

- In developer tools, under **sources** tab
- When running in developer mode source mapping is used to link bundled code to the original source files
- Under `localhost:4200`, `main.js` - clicking any line number will open source file
- All source files are also located in the `webpack://` section
- Can add breakpoints etc
- Chrome extension _Augury_ can be added to dev tools 

## Binding to Custom Properties

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
  @Input() childName: string; // decorator makes this a 'prop' 
}
```

In the `parent` component, where `nameFromParent` is available, the `childName` field can be bound to `nameFromParent`:
```angular2html
<app-child [childName]="nameFromParent"></app-child>
```

Then,  in the `child` component `childName` will have the value of `nameFromParent`:

```angular2html
<div>{{ childName }}</div>
```

 


 





 
    



 


 

 
  



