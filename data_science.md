### Data Literacy

Data Quality & Ethics
- High data quality leads to better decision-making, especially in healthcare.
- Addressing bias strengthens data reliability and reveals deeper truths.

---

Statistics in Legal & Social Issues
- Statistics can tackle large-scale, complex issues (e.g., hiring discrimination).
- Quantitative analysis helps in understanding systemic problems beyond individual cases.

---

Data Visualization
- Visuals are powerful tools for communicating data insights.
- The Challenger shuttle case highlights the risks of poor visualization.
- Effective data viz aids meaningful and accurate decision-making.

---

Data Analysis
- Context matters: understand both the numbers and their meaning.
- Recognizing the difference between correlation and causation is crucial.
- Example: Modern medicine benefits from causal analysis and smart data visualization.

<br>

**Key Takeaway**

Data literacy empowers individuals to ask better questions, make informed decisions, and approach complex issues with clarity and confidence.

<br>
<br>
<br>


### Variable Types

Numerical Variables
- Represent **measured** or **counted** values.
- **Must include units** (e.g., feet, dollars, minutes).
- Two subtypes:
  - **Discrete**: Whole numbers from counting.
    - Example: Number of branches.
  - **Continuous**: Can include decimals from measurement.
    - Example: `Height (ft)`

- Without units, numbers lack meaning (e.g., "3 what?").

---

Categorical Variables
- Represent **qualities or characteristics** using words or symbols.
- Three subtypes:

1. **Nominal Variables** (Names Only)
    - Categories with **no order** or ranking.
    - Example: Tree species — `London Plane`, `Honeylocust`, `Pin Oak`.

2. **Dichotomous Variables** (Two Options)
    - Only **two logical possibilities**: `Yes/No`, `True/False`, `0/1`.
    - Example: `Single` (Is the tree growing alone? Yes or No)
    - If there's a **third option**, it's **not** dichotomous.

3. **Ordinal Variables** (Ordered Categories)
    - Categories with a **natural order**, but not a true numeric scale.
    - Example: Tree prettiness rating (`1` = Least pretty, `5` = Most pretty).
    - Often used in:
        - **Likert scales**: "Strongly Disagree" → "Strongly Agree"
        - **Ranking**: 1st, 2nd, 3rd place in a competition.
    - Numbers represent **rank**, not quantity. You can’t average them meaningfully.

<br>

**Key Takeaways**
- **Numerical = Measurable/Countable** (with units).
- **Categorical = Descriptive** (can be named, binary, or ordered).
- **Ordinal variables may look numerical but are still categorical.**
- Always pair numerical data with **units** to clarify meaning.

<br>
<br>
<br>

### Missing Data Types

When working with datasets, it's common to encounter **missing values**. Understanding the **type of missingness** helps determine how to treat them.

---

**Missing Completely at Random (MCAR)**
- **Definition**: Data is missing for **no specific reason**.
- Example: You were cold, tired, or distracted during data collection.
- No pattern or underlying factor causes the missingness.
- **Key Point**: MCAR is not related to the data itself.

---

**Missing at Random (MAR)**
- **Definition**: The **missingness can be predicted** using other variables.
- Example: All `Redwood` trees are missing `Height` values.
- While not random in the usual sense, it is statistically "at random" **conditioned on other variables**.
- There may be a logical explanation — e.g., tape measures were too short for tall trees.
- **Key Point**: Requires deeper analysis and domain knowledge.

---

**Structurally Missing**
- **Definition**: The data **isn’t supposed to be there** in certain cases.
- Example: Counting visible fruit on trees. If a tree has no fruit, there's no count to be made.
- **Key Point**: These values are **intentionally** missing.

---

What To Do About Missing Data?

- **Structurally Missing**: Can be safely ignored — we don’t expect a value there.
- **MCAR / MAR**: Requires careful handling.
- **Any action** (even doing nothing) will **affect your analysis**.  

