# sg

## Local Application Setup 
1. Clone project: `git clone https://github.com/stevenelleman/sg.git`
2. `cd ./sg/app`
3. `yarn start`
4. Navigate to `http://localhost:3000/`
5. Additional information can be found in the [app docs](./app/README.md)

## Docs
- Spec ([./docs/spec.md](./docs/spec.md))
- App Setup Docs ([./app/README.md](./app/README.md))
- Frontend + API Plan ([./app/src/components/README.md](./app/src/components/README.md))

## Code-Amendment Best Practices 
- Don't commit directly to master, make PRs. 
- PRs should be reviewed and be required to have a +1 from non-contributor to the PR.
- PRs should have corresponding tests. 
- PRs should include documentation. Documentation should be kept with its relevant code to minimize drift.
- Correct known trip-wires (little annoyances, mistakes, gotchas) immediately - quick fixes pay off dividends.
- Prefix branch name with creator's initials.

## Tasks 
- [ ] Deployment Process [Tuesday]
- [ ] Local VM Setup [Friday]
- [ ] Frontend w/ Mock Endpoint/Data [Tuesday]
- [ ] Backend w/ Mock Data [Wednesday]
- [ ] Postgres DB [Thursday]