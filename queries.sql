-- 1) Find all current managers of each department and display his/her title, first name, last name, current salary.
SELECT  d.dept_name, t.title, e.first_name, e.last_name, s.salary
FROM employees e
RIGHT JOIN dept_manager dm
ON e.emp_no = dm.emp_no
LEFT JOIN departments d
ON d.dept_no = dm.dept_no
LEFT JOIN salaries s
ON e.emp_no = s.emp_no
LEFT JOIN titles t
ON e.emp_no = t.emp_no
WHERE s.to_date > NOW()
AND  dm.to_date > NOW()
AND t.to_date > NOW()
;

-- 2 Find all employees (department, title, first name, last name, hire date, how many years they have been working) to congratulate them on their hire anniversary this month.
SELECT  d.dept_name, t.title, e.first_name, e.last_name, e.hire_date, YEAR(NOW()) - YEAR(e.hire_date) AS Exp
FROM employees e
RIGHT JOIN dept_emp de
ON de.emp_no = e.emp_no
LEFT JOIN departments d
ON d.dept_no = de.dept_no
LEFT JOIN salaries s
ON e.emp_no = s.emp_no
LEFT JOIN titles t
ON e.emp_no = t.emp_no
WHERE de.to_date > NOW()
AND t.to_date > NOW()
AND s.to_date > NOW()
AND MONTH(e.hire_date) = MONTH(NOW())
AND YEAR(NOW()) - YEAR(e.hire_date) > 1
;

-- 3 Find all departments, their current employee count, their current sum salary.
SELECT d.dept_name, count(de.emp_no) AS employee_count, sum(s.salary) AS sum_salary
FROM departments d 
RIGHT JOIN dept_emp de
ON de.dept_no = d.dept_no
RIGHT JOIN employees e
ON de.emp_no = e.emp_no
RIGHT JOIN salaries s 
ON e.emp_no = s.emp_no
WHERE de.to_date > NOW()
AND s.to_date > NOW()
GROUP BY d.dept_no
; 



