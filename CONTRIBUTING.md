
## How to Contribute New Code

We welcome your contributions! Please follow the steps below to ensure consistency and CI compatibility:

### 1️⃣ Test your changes locally

Before submitting any code, make sure it works as expected by testing it locally.

Follow the steps in our [Development Guide](https://github.com/valkiriaaquatica/provider-proxmox-bpg?tab=readme-ov-file#developing) to run and test the provider on your machine.

---

### 2️⃣ Create a new branch

Use a meaningful name and follow the [Conventional Commit](https://www.conventionalcommits.org/) format for branches:

```bash
# For a bug fix
git checkout -b fix(my-feature):short-description

# For a new feature
git checkout -b feat(my-feature):short-description
```

> 📝 Use only lowercase for `fix`, `feat`, etc. Keep the branch name short and descriptive.

---

### 3️⃣ Make your changes and commit properly

Follow the **Conventional Commits** standard when writing your commit messages:

```bash
git commit -m "fix(my-feature): fix crash when API token is empty"
```

**Examples:**

* `feat(api): add support for token-based auth`
* `fix(vm): correct VM clone logic on large disks`
* `docs(readme): clarify dev setup instructions`

> ℹ️ Commits that don’t follow the Conventional Commit standard will **fail CI checks**.