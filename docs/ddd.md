# Nuky: Development Design Document

### Scheduled Jobs

1. nukeState() / 24h 12:00 UTC:

---

### Functional Requirements

#### auth

```
   login()
   logout()
   register()
```

#### province:

```
   both functions below requires can_attack true
   every move should be checked by recaptcha

   getAllProvinces()

   attack()
   support()
```

#### game:

```
   getGame()
```

---

### Non Functional Requiremets

1. latency: 200ms

2. rps: 1/2 x 5 x 100 = 250 RPS

3. memory: 100mb

4. cpu: graviton-small 60%

---

### Entities

#### user

```js
{
   username: string, UNIQUE
   email: string, UNIQUE
   lastMoveDate: Date

   password: hashed string

   updatedDate: Date
   deletedDate: Date
}
```

#### province

```js
{
    id: int, UNIQUE
    provinceName: string
    provinceColorHex: string
    attackCount: int
    supportCount: int
    destroymentRound: int // -1 = notDestroyed

    updatedDate: Date
    deletedDate: Date
}
```

### Development flow

type -> dto -> hook -> components -> view -> index.ts
