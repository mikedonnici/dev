# Vue.js and Nuxt.js

## Vuex

### State and Getters (and Computed Properties)

It's good practice to access store `state` and `getters` from `computed` properties.

```vue
<template>
    <div>
        {{ foo }} - {{ fooCount }}
    </div>
</template>

<script>
export default {
    computed: {
        foo() {
            return this.$store.state.foo
        },
        fooCount() {
            return this.$store.getters.fooCount
        }
    }
}
</script>
```

To avoid having to set up a computed property for each state item or getter, can use [`mapState`](https://vuex.vuejs.org/guide/state.html#the-mapstate-helper) and [`mapGetters`](https://vuex.vuejs.org/guide/getters.html#the-mapgetters-helper) helpers provided by `vuex`:

```vue
<template>
    <div>
        {{ foo }} - {{ fooCount }}
    </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'

export default {
    computed: {
        ...mapState(['foo']),
        ...mapGetters(['fooCount'])
    }
}
</script>
```

The args are arrays of state/getter field names required for this component, and the `computed` properties created will have the same name.

### Dynamic Getters

Usually, a computed property does not accept an argument because it is designed to behave more like a variable:

```vue
<template>
    <div>
        {{ computedFoo }}
    </div>
</template>

<script>
export default {
    computed: {
        computedFoo() {
            return "Bar"
        }
    }
}
</script>
```

However, if the `computed` property returns a function then it can receive an argument:

```vue
<template>
    <div>
        {{ computedFoo("Bar") }}
    </div>
</template>

<script>
export default {
    computed: {
        computedFoo() {
            return (str) => {
                return str
            }
        }
    }
}
</script>
```

This provides a way to pass an argument to a getter:

```vue
<template>
    <div>
        {{ thingById(123) }}
    </div>
</template>

<script>
export default {
    computed: {
        thingById() {
            return this.$store.getters.getThingById // <- returns a func
        }
    }
}
```

In the store:

```js
getters: {
    getThingById(state) {
        return (id) = {
            return state.things.find((thing) => {
                return thing.id === id
            })
        }
    }
}
```

Clever!!

Ref: <https://www.vuemastery.com/courses/mastering-vuex/vuex-state-getters>

## Mutations & Actions

### Committing Mutations (synchronous)

In a store, `mutations` are used to alter state.

It is common practice to name `mutations` in CAPITAL_SNAKE_CASE.

```js
state: {
    count: 0
}
mutations: {
    INC_COUNT(state) {
        state.count++
    }
}
```

In a component `mutations` are invoked using `commit`:

```vue
<script>
export default {
    methods: {
        incrementCount() {
            this.$store.commit('INC_COUNT')
        }
    }
}
</script>
```

A mutation can also receive a payload, in this case `value`:

```js
state: {
    count: 0
}
mutations: {
    INC_COUNT(state, value) {
        state.count += value
    }
}
```

```vue
<script>
export default {
    data() {
        return {
            incrementValue: 4
        }
    },
    methods: {
        incrementCount() {
            this.$store.commit('INC_COUNT', this.incrementValue)
        }
    }
}
</script>
```

### Dispatching Actions (asynchronous)

Actions can wrap business logic around mutations.

**Should _always_ put `mutations` within `actions`** to future-proof the app. This will avoid the need for heavy refactoring if and when the businss logic around the mutation is added or changes.

In the store, `actions` receive a context object, and an optional payload:

```js
actions: {
    someAction(ctx, payload) {
        // ...
    }
}
```

Object destructuring is often used for the context object to get the required items:

```js
actions: {
    someAction({state, commit}, payload) {
        // ...
    }
}
```

An action is _dispatched_ from a vue component:

```vue
<script>
export default {
    methods: {
        doThing(payload) {
            this.$store.dispatch('someAction', payload)
        }
    }
}
</script>
```

## Modules

Vuex can be arranged as modules to make larger stores easier to organise.

There are two main ways this is generally done:

### 1. Exporting constants

Benefit is private variables and methods.

`/store/store.js`

```js
import * as name from '@/store/modules/name.js'
```

`/store/modules/name.js`

```js
export const state = {...}
export const mutations = {...}
export const actions = {...}
export const getters = {...}
```

### 2. Exporting one object

```js
import name from '@/store/modules/name.js'
```

```js
export default {
    state = {...},
    mutations = {...},
    actions = {...},
    getters = {...}
```

### rootState

The `context` object that is passed to an action has `rootState`.

This provides a convenient way to access state across different modules:

```js
export const actions = {
    createFoo( { commit, rootState }, event) {
        console.log(rootState.moduleName.fieldName)
    }
}
```

### Actions, Mutations and Getters in the Global Namespace

The `state` of other modules can be accessed using `rootState` however, `actions`, `mutations` and `getters` are _always_ registered under the global namespace - ie `$store`.

So, regardless of where they are declared they are called without their module name.

However, this can very easily result in name collisions. Hence the use of _module namespacing_.

### Module Namespacing

Namespacing is implemented to avoid naming collisions and can be switched on for a particlar modules, eg:

`/store/modules/foo.js`

```js
export const namespaced = true
export const state = {...}
export const mutations = {...}
export const actions = {...}
export const getters = {...}
```

With this set, all `mutations`, `actions` and `getters` will be namespaced under 'foo', eg. `$store.dispatch('foo/action', payload)`

### mapActions, mapGetters Helper with Module Namespacing

To map actions to local methods, two syntaxes can be used.

Specify each action under its namespace:

```vue
<script>
import { mapActions } from 'vuex'
export default {
    methods: mapActions(['module/actionOne', 'module/actionTwo'])
}
</script>
```

...or pass the namespace as first argument to `mapActions()`

```vue
<script>
import { mapActions } from 'vuex'
export default {
    methods: mapActions('module', ['actionOne', 'actionTwo'])
}
</script>
```

The same syntax can be applied to `mapGetters()`

### Calling Namespaced Actions in Other Modules

Namespacing syntax is not required when calling `actions` or `getters` from within the same module. In addition, `mutations` from other should not be called directly. Instead, `actions` should be called that _mutate_ the state.

To call an action in another module:

```js
dispatch('module/action', payload, {root: true})
```

Where `payload` can be `null` as must specify `{root: true}` to look for the action from the root of the store.


## Success and Error Notifications









## References

[Vue Mastery](https://www.vuemastery.com/courses/)