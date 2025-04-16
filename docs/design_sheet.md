# Nuky: Development Design Sheet

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

#### move

```
   both functions below requires can_attack true
   every move should be checked by recaptcha

   attack()
   support()
```

#### state:

```
   getStates()
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
   email: string, UNIQUE
   password: hashed string
   last_move_date: Date
}
```

#### province

```js
{
    id: int, UNIQUE
    province_name: string
    province_color_hex: string
    attack_count: int
    support_count: int
}
```

#### game

```js
{
   loser_provinces: Province[]
}
```
