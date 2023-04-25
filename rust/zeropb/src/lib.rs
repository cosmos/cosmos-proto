extern crate core;
extern crate alloc;

use core::marker::PhantomData;
use core::ops::{Deref, DerefMut};
use alloc::alloc::{alloc_zeroed, dealloc, Layout};

pub struct Root<T> {
    _marker: PhantomData<T>,
    data: *mut u8,
    size: usize,
    capacity: usize,
}

impl<T> Root<T> {
    pub fn new() -> Self {
        const DEFAULT_CAP: usize = 65536;
        Self {
            _marker: PhantomData,
            data: unsafe { alloc_zeroed(Layout::from_size_align_unchecked(DEFAULT_CAP, 8)) },
            size: core::mem::size_of::<T>(),
            capacity: DEFAULT_CAP,
        }
    }
}

impl<T> Deref for Root<T> {
    type Target = T;

    fn deref(&self) -> &Self::Target {
        unsafe { &*(self.data as *const T) }
    }
}

impl<T> DerefMut for Root<T> {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut T) }
    }
}

impl<T> Drop for Root<T> {
    fn drop(&mut self) {
        unsafe {
            dealloc(self.data, Layout::from_size_align_unchecked(self.capacity, 1))
        }
    }
}