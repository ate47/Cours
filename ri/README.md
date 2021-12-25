- [Practicals](#practicals)
- [Projects](#projects)
- [Measures](#measures)
  - [Measures for unranked sets](#measures-for-unranked-sets)
  - [Measures for ranked lists](#measures-for-ranked-lists)

# Practicals

- [TP1](https://claroline-connect.univ-st-etienne.fr/web/app.php/resource/open/1386986)

# Projects

- [Repo of the project](https://github.com/ThomasGUICHARD/RechercheInformation)

# Measures

## Measures for unranked sets

Get a measure for a unordered set of results.

|               | Relevant | Non Relevant |
| ------------- | -------- | ------------ |
| Retrieved     | TP       | FP           |
| Non Retrieved | FN       | TN           |

- Precision: $$P = \frac{TP}{TP + FP}$$
- Recall: $$R = \frac{TP}{TP + FN}$$
  Recall can be maximise by sending all the documents (precision will be near $0$, but the recall would be $1$)
- Accuracy: $$A = \frac{TP + TN}{FP + FN}$$
  In IR, accuracy can be maximise by sending 0 document, accuracy will be close to $0.99$, but the precision and the recall will also be $0$.
- F measure deal with that
  $$F = \frac{1}{\alpha \times (1 - \alpha)\times \frac{1}{R}}$$
  we can set
  $$\beta^2 = \frac{1 - \alpha}{\alpha}$$
  to replace
  $$F = \frac{(1 + \beta^2) \times P \times R}{\beta^2 \times P + R}$$
  usually we take $\beta = 1$ or $\alpha = 0.5$.
  $$F_1 = F_{\beta = 1} = \frac{2 \times P \times R}{P + R}$$

## Measures for ranked lists

Get measures for ordered sets, so we can have a $top_1, top_2, \dots top_k$ result list.

We can recreate the the previous measure with $k$ result(s), so we have the Presision@k, Recall@k, etc. We only count the $k$ first values with their ranks.
