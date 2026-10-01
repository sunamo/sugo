---
schema_version: 6
type: library
file_count: 4
avg_lines_per_file: 56
move_to_legacy_percent: 30
generated_date: 2026-10-01
generated_time: 16:40:46
github_source_url: 
last_build_ok: 
last_build_date: 
last_tests_run_date: 
covered_lines: 
total_lines: 
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
