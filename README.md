
# Alif-task-payments  

## Task
Implement the Installment Payments service for a financial institution. This institution has just entered into an agreement with an online electronics store. Depending on the product category, the duration of the installment plan may be different. Installment range: 3-6-9-12-18-24 months. There are only three categories in the online image store with the following installment duration: Smartphone: installment [3-9] months Computer: installment [3-12] months TV: installment [3-18] months For each additional unit of range, 3% will be added for smartphones, 4% for computers and 5% for televisions of the total product. For example, if a customer wants to buy a smartphone with an installment plan for 18 months for 1,000 somoni, the total amount will be 1,060 somoni (1,000 somoni + 6%). Write a function/method where an e-shop sends four parameters to get the total payment amount for a product: product, amount, buyer's phone number, and installment range. Also, after each payment process, an SMS with purchase details will be sent to the customer's phone number.

## Run Locally  
Clone the project  

~~~bash  
  git clone https://github.com/ProninIgorr/alif-task-payments
~~~

Go to the project directory  

~~~bash  
  cd my-project
~~~



Start the app  

~~~bash  
./run.sh
~~~  
