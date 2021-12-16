from sqlalchemy import Column, Integer, String, Date
from sqlalchemy.orm import declarative_base

Base = declarative_base()

class Draw(Base):
    __tablename__ == 'draws'

    id = Column(Integer, primary_key=True)
    number = Column(String)
    date = Column(Date)
    balls = Column(String)
    stars = Column(String)
    order = Column(String)