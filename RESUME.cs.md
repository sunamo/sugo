---
schema_version: 11
type: my-library
category_override: Go_Projects
file_count: 4
file_extensions: md:2, go:1, noext:1
file_extensions_updated: 2026-10-04
avg_lines_per_file: 56
total_lines: 121
metrics_lm: 2026-10-01 16:40:46
move_to_legacy_percent: 30
description_updated: 2026-10-01
links_updated: 2026-10-01
github_source_url: not found
origin_status: found
origin_checked: 2026-10-01
article_source_url: not run
article_status: pending
article_checked: not run
last_build_ok: not run
last_build_date: not run
last_tests_run_date: not run
covered_lines: 0
---

## Description

Vlastní sdílená knihovna pomocných funkcí pro Go (balíček `sugo`).
V jediném souboru `sugo.go` jsou funkce jako `Join`, `ToStr`, `FirstChar`, `Contains`, `PrintFactors`, `Int`/`Int64` (konverze polí) a `Exc` pro výpis chyb.
Je to malá osobní utilitka bez testů.

## Původ zdrojáků

Staženo z GitHubu: **ne** — vlastní knihovna autora (`sunamo/sugo`)

- Ověřeno: `gh api repos/sunamo/sugo` vrací `fork: false` bez rodiče; README "My shared library for Go"; historie 16 commitů jen od uživatele (Radek Jancik / Radek Jančík / VPS); v kódu žádné URL ani cizí copyright.

## Doporučení přesunu do legacy

Doporučení přesunu do sunamocz-legacy.visualstudio.com: **30 %** — vlastní sdílená knihovna, ale malá

- `sunamo/sugo` je autorova vlastní knihovna (16 commitů), nejde o cizí kód.
- Jen 4 soubory, proto nižší riziko při smazání, ale obsah je unikátní, takže nejde o vysoké číslo.

## Vazby na moje repa

- Submoduly: žádné
- ProjectReference / PackageReference: žádné
